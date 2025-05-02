FROM ubuntu:22.04

ENV RUNNER_VERSION=2.317.0 \
    DEBIAN_FRONTEND=noninteractive \
    RUNNER_USER=runner


# Instala dependências
RUN apt-get update && \
    apt-get install -y curl unzip tar git sudo libicu70 libssl3 ca-certificates jq && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Cria diretório e usuário
RUN useradd -m ${RUNNER_USER}
WORKDIR /home/${RUNNER_USER}

# Baixa o runner do GitHub
RUN curl -o actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz -L https://github.com/actions/runner/releases/download/v${RUNNER_VERSION}/actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz && \
    tar xzf actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz && \
    rm actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz && \
    chown -R ${RUNNER_USER}:${RUNNER_USER} .

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    apt-get update && apt-get install -y unzip && \
    unzip awscliv2.zip && \
    ./aws/install && \
    rm -rf aws awscliv2.zip

# Copia o script de entrada
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Usa usuário não-root
USER ${RUNNER_USER}

# Define ponto de entrada
ENTRYPOINT ["/entrypoint.sh"]