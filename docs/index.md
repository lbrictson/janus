# Welcome to Janus

Janus is a simple script execution program contained in a single binary. It is designed as a portal to allow non-technical
users to run scripts authored by more technical users.  Both adhoc and scheduled job execution is supported.

## Key Features

- **Simple**: single binary that can be run on any system with no dependencies.
- **Job Notifications**: receive notifications when jobs start, run and complete.  Supported channels
    - Email
    - Slack
    - Discord
    - Webhook
    - SMS (via Twilio)
    - PagerDuty
    - AWS SNS
    - AWS EventBridge
- **Job Scheduling**: schedule jobs to run at specific times or intervals using cron syntax.
- **Job Parameters**: pass parameters to jobs at runtime.
- **Job Output**: view job output in the UI.
- **Job History**: view job history in the UI.
- **Job versioning**: version jobs to allow for changes over time.
- **Permissions**: group jobs into projects and set project permissions on a per user basis
- **REST API**: interact with Janus via a REST API
- **Webhooks**: trigger jobs via webhooks
- **Secrets**: store secrets in Janus and reference them in jobs
- **Storage**: supports both sqlite and postgres databases