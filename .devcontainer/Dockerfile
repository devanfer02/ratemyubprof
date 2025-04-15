FROM golang:1.24-bookworm

# Install curl, gnupg, etc.
RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

# Add Docker’s official GPG key
RUN install -m 0755 -d /etc/apt/keyrings && \
    curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg && \
    chmod a+r /etc/apt/keyrings/docker.gpg

# Add the Docker APT repo
RUN echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null


RUN apt-get update && apt-get install -y \
    git \
    vim \ 
    libpq-dev \
    docker-ce-cli \ 
    iputils-ping \ 
    && rm -rf /var/lib/apt/lists/*

# Install air live reload golang project
RUN go install github.com/air-verse/air@latest

# Install go-migrate cli
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.linux-amd64.tar.gz | tar xvz  && \
mv migrate /go/bin/migrate && \
chmod +x /go/bin/migrate