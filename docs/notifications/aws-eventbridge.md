# AWS Event Bridge

Janus is capable of putting events onto the default or custom AWS EventBridge buses.  This allows you to trigger AWS services based on Janus job events.

Create a new AWS EventBridge channel at `/notifications`

## Configuration

| Config Item       | Explanation                       |
|-------------------|-----------------------------------|
| Event Bus Name    | Name of the event bus             |
| Event Source      | Source of the event               |
| Detail Type       | Detail type for this custom event |
| AWS Region        | AWS Region for this event bridge  |
| Access Key ID     | AWS Access key ID                 |
| Secret Access Key | AWS Secret access key             |

## Event Source

Event source is a great way to filter events in AWS EventBridge.  Configuring a unique name for your environment will allow you to easily filter
for Janus events.

Read more here:  https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_EventSource.html

## Detail Type

Detail type is a way to filter events in AWS EventBridge.  Configuring a unique name for your environment will allow you to easily filter

Read more here:  https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutEvents.html

## Permissions

The AWS Access Key ID and Secret Access Key should be for an IAM user with the following permissions:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "events:PutEvents"
            ],
            "Resource": "*"
        }
    ]
}
```
