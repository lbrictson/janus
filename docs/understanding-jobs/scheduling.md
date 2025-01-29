# Scheduling

Janus allows jobs to be scheduled to run on a fixed schedule.  Schedules are always defined in cron syntax.

Read more about cron syntax [here](https://crontab.guru/).

## Creating a Schedule

When editing or creating a job set a valid cron expression in the "Cron Schedule" field and flip the "Schedule Enabled" slider to the on position.

Jobs will fail to save if the schedule is invalid.

## Editing Behavior

The following cases have these outcomes:

- Job has schedule enabled and you change it to not:  The job will no longer run on the schedule.
- Job has schedule disabled and you enable it:  The job will run on the schedule.
- Job has schedule enabled and you change the schedule:  The job will run on the new schedule.