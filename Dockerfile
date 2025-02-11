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
RUN apt update && apt install -y ca-certificates python3 python3-pip curl wget pipx awscli ssh sshpass git gnupg apt-transport-https
RUN pipx install --include-deps ansible && pipx install --include-deps pipenv && pipx ensurepath
RUN curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.32/deb/Release.key | gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg && chmod 644 /etc/apt/keyrings/kubernetes-apt-keyring.gpg
RUN echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.32/deb/ /' | tee /etc/apt/sources.list.d/kubernetes.list && chmod 644 /etc/apt/sources.list.d/kubernetes.list
RUN apt update && apt install -y kubectl

EXPOSE 8080
EXPOSE 8081
ENV JANUS_DB_PATH=/data/janus.db

ENTRYPOINT ["/janus"]