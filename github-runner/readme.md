




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
    -e GITHUB_URL=https://github.com/lucasirc/github-actions-testing \
    -e RUNNER_TOKEN=${GITHUB_RUNNER_TOKEN} \
    -e RUNNER_NAME=gh-testing-runner  \
    -e RUNNER_LABELS=digiworld,ubuntu \
    github-runner
  
```