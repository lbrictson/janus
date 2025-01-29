# AWS SNS

Janus is capable of sending notifications to AWS SNS topics.  This allows you to trigger AWS services based on Janus job events or receive
notifications via email, sms and other channels.

Create a new AWS SNS channel at `/notifications`

## Configuration

| Config Item       | Explanation                       |
|-------------------|-----------------------------------|
| Topic ARN         | ARN of the SNS topic              |
| AWS Region        | AWS Region for this SNS topic     |
| Access Key ID     | AWS Access key ID                 |
| Secret Access Key | AWS Secret access key             |

## Permissions

The AWS Access Key ID and Secret Access Key should be for an IAM user with the following permissions:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "sns:Publish"
            ],
            "Resource": "$yourArnGoesHere"
        }
    ]
}
```

## Notes

The topic must already exist, Janus does not attempt to create the topic if it is missing.