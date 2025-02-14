# janus

Janus is a simple Job (script) runner that is designed to be easy to use and easy to deploy.  The general purpose is to 
allow non-technical users to execute scripts authored by more technical users without needing to update their local or enviroment
or get direct access to various systems.

![image](https://github.com/user-attachments/assets/bc81f684-cf76-43c4-8499-a56fd9853551)

![image](https://github.com/user-attachments/assets/f109e591-9747-4db9-8051-432da3d61f1c)

![image](https://github.com/user-attachments/assets/4c72977c-0794-44e6-b0d6-a60732812368)

![image](https://github.com/user-attachments/assets/b59c73c0-9417-420f-9c11-7dac169af5ea)

![image](https://github.com/user-attachments/assets/9f8cd912-b7a2-44f1-958a-91c69a51dcf9)


## Features

- Single binary deployment
- Simple web interface
- Project based permissions
- Notification system
- Job versioning
- Secret Storage
- Scheduled/Recurring Jobs
- Webhook Trigger Support
- Single Sign On Support
- SQLite and PostgreSQL Support

Read the docs:  [Documentation](https://janus.brictson.dev)

or jump right to the quick start which only takes a few seconds to get running:  [Quick Start](https://janus.brictson.dev/getting-started/installation/)


## Local Development

### Prerequisites

- Go
- Docker (optional)
- Air (optional for hot reloading)

### Running tests

```bash
go test ./...
```

### Running the server
```bash
go run cmd/server/main.go
# Or if you have air installed simply `air`
# Access at http://localhost:8080
# Username: admin@localhost
# Password: ChangeMeBeforeUse1234!
```

### Adding Database models

```bash
go run -mod=mod entgo.io/ent/cmd/ent new $SCHEMA-NAME
```

### Docker Development

You can use the provided docker-compose file to run a simple mail server to test email notifications.

```bash
docker-compose up -d
# Access at http://localhost:8025
# SMTP is at localhost:1025, username and password can be anything
```
