package main

import (
	"fmt"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func getSubnetsByVPC(ctx *pulumi.Context, vpcId string) ([]string, error) {
	subnets, err := ec2.GetSubnets(ctx, &ec2.GetSubnetsArgs{
		Filters: []ec2.GetSubnetsFilter{
			{
				Name:   "vpc-id",
				Values: []string{vpcId},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return subnets.Ids, nil
}

func main() {
	repoName := "lucasirc/github-actions-testing"
	awsAccountId := os.Getenv("AWS_ACCOUNT_ID")
	ecrUrl := fmt.Sprintf("%s.dkr.ecr.us-east-1.amazonaws.com/gh-runner:latest", awsAccountId)

	GITHUB_RUNNER_TOKEN := os.Getenv("GITHUB_RUNNER_TOKEN")
	if GITHUB_RUNNER_TOKEN == "" {
		fmt.Println("GITHUB_RUNNER_TOKEN: env variable is required.")
		os.Exit(1)
	}

	pulumi.Run(func(ctx *pulumi.Context) error {

		// VPC padrão
		vpc, err := ec2.LookupVpc(ctx, &ec2.LookupVpcArgs{Default: pulumi.BoolRef(true)})
		if err != nil {
			return err
		}

		subnets, err := getSubnetsByVPC(ctx, vpc.Id)
		if err != nil {
			return err
		}

		// ECS Cluster
		cluster, err := ecs.NewCluster(ctx, "github-runner-cluster", nil)
		if err != nil {
			return err
		}

		// IAM Role para execução da Task
		execRole, err := iam.NewRole(ctx, "task-exec-role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(`{
			  "Version": "2012-10-17",
			  "Statement": [{
			    "Effect": "Allow",
			    "Principal": {
			      "Service": "ecs-tasks.amazonaws.com"
			    },
			    "Action": "sts:AssumeRole"
			  }]
			}`),
		})
		if err != nil {
			return err
		}

		_, err = iam.NewRolePolicyAttachment(ctx, "task-exec-policy", &iam.RolePolicyAttachmentArgs{
			Role:      execRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"),
		})
		if err != nil {
			return err
		}

		// Task Definition
		taskDef, err := ecs.NewTaskDefinition(ctx, "github-runner-task", &ecs.TaskDefinitionArgs{
			Family:                  pulumi.String("github-runner"),
			Cpu:                     pulumi.String("512"),
			Memory:                  pulumi.String("1024"),
			NetworkMode:             pulumi.String("awsvpc"),
			RequiresCompatibilities: pulumi.StringArray{pulumi.String("FARGATE")},
			ExecutionRoleArn:        execRole.Arn,
			ContainerDefinitions: pulumi.String(fmt.Sprintf(`[{
				"name": "runner",
				"image": "%s",
				"essential": true,
				"environment": [
					{"name": "GITHUB_URL", "value": "https://github.com/%s"},
					{"name": "RUNNER_TOKEN", "value": "%s"},
					{"name": "RUNNER_NAME", "value": "gh-ecs-runner"},
					{"name": "RUNNER_LABELS", "value": "ecs,fargate,linux,digiworld"}
				]
			}]`, ecrUrl, repoName, GITHUB_RUNNER_TOKEN)),
		})
		if err != nil {
			return err
		}

		// Service
		_, err = ecs.NewService(ctx, "github-runner-service", &ecs.ServiceArgs{
			Cluster:        cluster.ID(),
			DesiredCount:   pulumi.Int(1),
			LaunchType:     pulumi.String("FARGATE"),
			TaskDefinition: taskDef.Arn,
			NetworkConfiguration: &ecs.ServiceNetworkConfigurationArgs{
				Subnets:        pulumi.ToStringArray(subnets),
				AssignPublicIp: pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}

		return nil
	})
}
