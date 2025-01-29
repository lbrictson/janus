# Creating Your First Job

Jobs in Janus are the core of the system.  Jobs can be designed to run as one-shot tasks or as scheduled recurring tasks.

## Creating a Job

Creating a job requires a project, create one on the dashboard if need be first.

1. Navigate to the project you want to create a job in.
2. Click the `Create Job` button.

## Job Configuration

For your first job we will create a simple job runs a simple command called Say Hello.

Name your job "Say Hello" and give it a fun description.

Add an argument to the job called "NAME" and leave the default value blank.

Add another argument to the job called "MESSAGE" and set the default value to "Hello", under allowed values add "Hello", "Goodbye", "Howdy".

For the script Enter something that looks like this
```bash
#!/bin/bash
echo "{{JANUS_ARG_MESSAGE}}, {{JANUS_ARG_NAME}}!"
sleep 5
echo "Goodbye, $JANUS_ARG_NAME"
```
This script is quite simple, but it gives a good introduction to how to interact with arguments in Janus.

Arguments will always be uppercased and prefixed with `JANUS_ARG_`.  They can be interpolated into the script using the double curly brace syntax.

All arguments are also present in the environment as well, so you can use them in the script like you would any other environment variable.

The script above demonstrates using both the environment variable and the interpolated syntax.

When checking the "Sensitive" box, the argument will be hidden from the logs and the UI.

Save your job and run it to see the output.
