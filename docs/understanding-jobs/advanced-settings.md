# Advanced Settings

There are some settings available that are fine on their defaults but can be changed as needed.

### Timeout (seconds)

The timeout setting is the maximum amount of time a job can run before it is killed.  The default is 3600 seconds (1 hour).  If you have a job that you know will take longer than an hour, you can increase this setting.

### Require File Upload

When enabled, this setting will require that a file be uploaded to the job before it can be run.  This is useful for jobs that require a file to be present before they can run.

The file is always saved into the jobs working directory as "file".  It is up to your script to rename it to the file type 
or parse it as is.  Allowing users to upload files can be a security risk, so be sure to validate the file before using it.

### Allow Concurrent Runs

When enabled, this setting will allow the job to run multiple times concurrently.  This is useful for jobs that can be run in parallel.

This also applies to scheduled runs and webhook triggers.  If a job is running and another run is triggered, the new run will start immediately.

