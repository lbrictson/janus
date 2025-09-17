# Webhooks (Inbound Triggers)

Webhooks provide a way to trigger Janus jobs from external systems via HTTP requests. Each webhook has a unique URL that can be called to execute a specific job.

## Creating a Webhook

To create an inbound webhook:
1. Navigate to the Webhooks section in the Janus UI
2. Select the job you want to trigger
3. Optionally enable API key authentication
4. Save the webhook to get your unique webhook URL

## Webhook URL Format

Once created, your webhook will have a URL in the following format:
```
https://your-janus-instance/webhook/execute/{unique-key}
```

## Authentication

Webhooks support optional API key authentication for added security.

### Without Authentication
If API key authentication is disabled, anyone with the webhook URL can trigger the job.

### With API Key Authentication
When API key authentication is enabled:
1. An API key will be generated when you create the webhook
2. Include the API key in the request header: `X-API-KEY: your-api-key`
3. Requests without the correct API key will receive a 401 Unauthorized response

Example curl command with authentication:
```bash
curl -X POST https://your-janus-instance/webhook/execute/abc123 \
  -H "X-API-KEY: your-api-key" \
  -H "Content-Type: application/json" \
  -d '{"message": "Hello from webhook"}'
```

## POST Request Payloads

When triggering a webhook with a POST request, the request body will be available to your job as an environment variable:

- **Variable Name**: `$JANUS_ARG_WEBHOOK_PAYLOAD`
- **Content**: The raw body of the POST request (typically JSON)

### Example: Processing Webhook Data

```bash
#!/bin/bash
# Your job script can access the webhook payload
echo "Received webhook payload:"
echo "$JANUS_ARG_WEBHOOK_PAYLOAD"

# Parse JSON payload (requires jq)
echo "$JANUS_ARG_WEBHOOK_PAYLOAD" | jq '.message'
```

### Example: GitHub Webhook Handler

```bash
#!/bin/bash
# Handle GitHub webhook events
EVENT_TYPE=$(echo "$JANUS_ARG_WEBHOOK_PAYLOAD" | jq -r '.action')
REPO_NAME=$(echo "$JANUS_ARG_WEBHOOK_PAYLOAD" | jq -r '.repository.name')

echo "GitHub event: $EVENT_TYPE on repository: $REPO_NAME"

case $EVENT_TYPE in
  "opened")
    echo "New PR opened"
    # Add your PR handling logic here
    ;;
  "closed")
    echo "PR closed"
    # Add your cleanup logic here
    ;;
esac
```

## Limitations

- **File Uploads**: Jobs that require file uploads cannot be triggered via webhooks
- **Arguments**: All job arguments must have default values, as webhooks cannot provide runtime argument values (except for the webhook payload itself)
- **Scheduled Jobs**: Jobs triggered by webhooks will not affect the job's regular schedule if one exists

## GET vs POST Requests

- **GET Requests**: Trigger the job with default argument values only
- **POST Requests**: Trigger the job with default argument values plus access to the request body via `$JANUS_ARG_WEBHOOK_PAYLOAD`

## Response Codes

| Code | Description |
|------|-------------|
| 200  | Job triggered successfully |
| 401  | Invalid or missing API key (when authentication is enabled) |
| 500  | Job failed to start (check job requirements) |

## Best Practices

1. **Use API Keys**: Always enable API key authentication for production webhooks
2. **Validate Payloads**: Add validation in your job scripts to handle unexpected payload formats
3. **Error Handling**: Include error handling in your scripts for malformed webhook data
4. **Logging**: Log webhook payloads for debugging and audit purposes
5. **Timeout Settings**: Configure appropriate timeout values for jobs triggered by webhooks