




# GitHub Runner

## build docker

```sh
cd github-runner
docker build --platform=linux/amd64 -t github-runner .
```




## run docker 

```sh

docker run -it --rm  --platform=linux/amd64   \ 
--name gh-runner   \ 
-e GITHUB_URL=https://github.com/lucasirc/github-actions-testing   \
-e RUNNER_TOKEN=${GITHUB_RUNNER_TOKEN}   \
-e RUNNER_NAME=gh-testing-runner  \
github-runner

```

##run docker

```sh

docker run -it --rm  --platform=linux/amd64 \
    --name gh-runner \
    -p 8080:8080 \
    -e GITHUB_URL=https://github.com/lucasirc/github-actions-testing \
    -e RUNNER_TOKEN=${GITHUB_RUNNER_TOKEN} \
    -e RUNNER_NAME=gh-testing-runner  \
    -e RUNNER_LABELS=digiworld,ubuntu \
    github-runner
  
```


## upload container to ECR


### login in ECR

```sh

aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.us-east-1.amazonaws.com

```

### create repository

```sh
aws ecr create-repository --repository-name gh-runner --region us-east-1
```

### create tag

```sh
docker tag github-runner:latest ${AWS_ACCOUNT_ID}.dkr.ecr.us-east-1.amazonaws.com/gh-runner:latest
```


### push image
```sh
docker push ${AWS_ACCOUNT_ID}.dkr.ecr.us-east-1.amazonaws.com/gh-runner:latest
```