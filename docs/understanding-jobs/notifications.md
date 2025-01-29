# Notifications

Job notifications are a great way to keep users informed without needing to visit Janus itself.  

Janus sends notifications optionally at three points in the lifecycle of a job:

1. When a job starts
2. When a job completes successfully
3. When a job fails

Each notification case above can have as many notification channels attached to it as you like.

## Notes

Failure notifications are sent for the following reasons:

1. The job script exits with a non-zero status code.
2. The job script times out.
3. Max concurrent runs is reached.
4. Job doesn't allow concurrent runs and another run is triggered while the job is already running.