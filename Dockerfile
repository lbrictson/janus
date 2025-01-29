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
RUN apt update && apt install -y ca-certificates python3 python3-pip curl wget pipx awscli ssh sshpass git
RUN pipx install --include-deps ansible && pipx install --include-deps pipenv && pipx ensurepath
EXPOSE 8080
EXPOSE 8081
ENV JANUS_DB_PATH=/data/janus.db

ENTRYPOINT ["/janus"]