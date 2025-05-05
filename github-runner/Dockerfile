FROM ubuntu:22.04

ENV RUNNER_VERSION=2.317.0 \
    DEBIAN_FRONTEND=noninteractive \
    RUNNER_USER=runner \
    RUNNER_GROUP=runner-group

RUN apt-get update && \
    apt-get install -y curl wget unzip tar git sudo libicu70 libssl3 ca-certificates jq nginx gosu && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN groupadd -r ${RUNNER_GROUP} && \
    useradd -r -u 1001 -g ${RUNNER_GROUP} -s /bin/bash -m ${RUNNER_USER} && \
    mkdir -p /app && \
    chown ${RUNNER_USER}:${RUNNER_GROUP} /app

WORKDIR /app
COPY --chown=${RUNNER_USER}:${RUNNER_GROUP} entrypoint.sh /app/entrypoint.sh


RUN curl -o actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz -L https://github.com/actions/runner/releases/download/v${RUNNER_VERSION}/actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz && \
    tar xzf actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz && \
    rm actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz


RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    ./aws/install && \
    rm -rf aws awscliv2.zip

RUN chmod +x /app/entrypoint.sh

# Define ponto de entrada
ENTRYPOINT ["/app/entrypoint.sh"]