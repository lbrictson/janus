# Arguments

Arguments are how you can accept user input, provide default values and restrict the values that can be passed to a job. Janus supports multiple input types to make data entry easier and more accurate.

## Input Types

Janus supports different input types for arguments to provide better user experience and data validation:

| Type | Description | UI Element | Example Use Case |
|------|-------------|------------|------------------|
| `string` | Text input (default) | Text field | Names, descriptions, paths |
| `number` | Numeric values only | Number input | Port numbers, counts, thresholds |
| `date` | Date selection | Date picker | Report dates, cutoff dates |
| `datetime` | Date and time selection | DateTime picker | Scheduled events, timestamps |

The input type affects how the argument is presented in the UI but all values are passed to your script as strings.

### Example: Using Different Input Types

```bash
#!/bin/bash
# Arguments with different types
echo "Name: $JANUS_ARG_NAME"           # string type
echo "Port: $JANUS_ARG_PORT"           # number type  
echo "Report Date: $JANUS_ARG_DATE"    # date type (format: YYYY-MM-DD)
echo "Start Time: $JANUS_ARG_START"    # datetime type (format: YYYY-MM-DDTHH:MM:SS)
```

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

### Scheduled Jobs
Scheduled jobs cannot have arguments without default values as there is no human to provide the value at runtime.

### Webhooks
Webhooks act in the same manner - all arguments must have default values. However, webhooks can pass data via POST request body which becomes available as `$JANUS_ARG_WEBHOOK_PAYLOAD`. See the [Webhooks documentation](webhooks.md) for more details.

### Input Type Considerations
- **Date/DateTime arguments**: Always passed to scripts in ISO 8601 format
- **Number arguments**: Validated in the UI but passed as strings to maintain compatibility
- **Dropdown with allowed values**: Works with any input type