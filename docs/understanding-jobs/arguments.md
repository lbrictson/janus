# Arguments

Arguments are how you can accept user input, provide default values and restrict the values that can be passed to a job.

## General Behavior

Arguments are always uppercased and prefixed with `JANUS_ARG_`.  They can be interpolated into the script using the double curly brace syntax.

All arguments are also present in the environment as well, so you can use them in the script like you would any other environment variable.

Assume you have a variable called `NAME`, this is how you would access it within your script.

```bash
#!/bin/bash
echo "Hello, $JANUS_ARG_NAME!"
echo "Hello, {{JANUS_ARG_NAME}}!"
```

## Default Values

Arguments can have default values.  If the argument is not provided by the user, the default value will be used.  When
starting the job from the UI the default value will be pre-filled if one was provided when creating the job.

## Allowed Values

Arguments can have allowed values.  If the user provides a value that is not in the allowed list, the job will fail to start.
When using allowed values the default value must be one of the allowed values.

## Sensitive

When checking the "Sensitive" box, the argument will be hidden from the logs and the UI when viewing history.

This is a great use case for passwords and API keys.

## Special Cases

Scheduled jobs cannot have arguments without default values as there is no human to provide the value at runtime

Webhooks act in the same manner.