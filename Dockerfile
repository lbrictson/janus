FROM golang:1.23 AS build-stage

WORKDIR /app

COPY . /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o /janus cmd/server/main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM debian:12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /janus /janus
RUN mkdir /data
RUN apt update && apt install -y ca-certificates python3 python3-pip curl wget pipx awscli ssh sshpass tree git gnupg apt-transport-https ansible jq yq rsync rclone software-properties-common gnupg httpie
RUN pipx install --include-deps ansible && pipx install --include-deps pipenv && pipx ensurepath
RUN curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.32/deb/Release.key | gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg && chmod 644 /etc/apt/keyrings/kubernetes-apt-keyring.gpg
RUN echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.32/deb/ /' | tee /etc/apt/sources.list.d/kubernetes.list && chmod 644 /etc/apt/sources.list.d/kubernetes.list
RUN wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor | tee /usr/share/keyrings/hashicorp-archive-keyring.gpg > /dev/null
RUN echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(grep -oP '(?<=UBUNTU_CODENAME=).*' /etc/os-release || lsb_release -cs) main" | tee /etc/apt/sources.list.d/hashicorp.list
RUN apt update && apt install -y kubectl terraform vault
RUN curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
RUN chmod 700 get_helm.sh && ./get_helm.sh && rm get_helm.sh

EXPOSE 8080
EXPOSE 8081
ENV JANUS_DB_PATH=/data/janus.db

ENTRYPOINT ["/janus"]