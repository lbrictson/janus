# Secrets

Secrets are always scoped to the project level, you cannot share secrets across projects.

## Creating/Editing a Secret

Inside the project there is a "Secrets" button along the top table border.  Secrets have a name and value.

## Using a Secret

Secrets are available to jobs as environment variables.  They are prefixed with `JANUS_SECRET_` and are always uppercased.

Example:  If you have a secret called `API_KEY`, you can access it in your script as `$JANUS_SECRET_API_KEY`.

```bash
#!/bin/bash
echo "Hello, $JANUS_SECRET_API_KEY!"
# Interpolated syntax is also available
echo "Hello, {{JANUS_SECRET_API_KEY}}!"
```

Secrets are always masked when being output into logs and the history UI.

