# Environment Variables

## General Behavior
Any Janus job that has arguments will have those arguments available as environment variables.  This allows you to use the arguments in your scripts as you would any other environment variable.

The variables are alaways uppercased and prefixed with `JANUS_ARG_`.

Example:  If you have an argument called `NAME`, you can access it in your script as `$JANUS_ARG_NAME`.

```bash
#!/bin/bash
echo "Hello, $JANUS_ARG_NAME!"
```

# Special Cases

## Webhook Payloads

When a job is triggered via an inbound webhook with a POST request, the request body is automatically made available as an environment variable:

- **Variable Name**: `JANUS_ARG_WEBHOOK_PAYLOAD`
- **Content**: The raw body of the POST request
- **Common Format**: Usually JSON, but can be any text format

This allows your job scripts to process data sent from external systems like GitHub, GitLab, CI/CD pipelines, or custom applications.

Example usage in a script:
```bash
#!/bin/bash
# Access the webhook payload
echo "Received payload: $JANUS_ARG_WEBHOOK_PAYLOAD"

# Parse JSON payload (requires jq)
EVENT_TYPE=$(echo "$JANUS_ARG_WEBHOOK_PAYLOAD" | jq -r '.event_type')
echo "Event type: $EVENT_TYPE"
```

For more details on setting up and using webhooks, see the [Webhooks documentation](webhooks.md).