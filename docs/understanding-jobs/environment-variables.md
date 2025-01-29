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

When receiving an incoming webhook to trigger a job the payload of the webhook (for POSTs) will be available to the job as
the environment variable `JANUS_ARG_WEBHOOK_PAYLOAD`.