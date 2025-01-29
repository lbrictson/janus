# Webhooks (Outbound)

Janus supports sending notifications as webhooks to external targets.  The webhook is JSON formatted and sent as a POST request.

## Configuration

| Config Item | Explanation                                                                                                                             |
|-------------|-----------------------------------------------------------------------------------------------------------------------------------------|
| Webhook URL | The URL of the webhook                                                                                                                  |
 | Headers | A list of headers to send with the request.  Headers are defined in JSON as pairs like: `{"x-api-key": "my api key", "user": "my-user"}` |

## Schema

```json
{
  "JobName": "TEST",
  "JobID": 0,
  "ProjectName": "TEST",
  "JobStatus": "Testing",
  "JobDurationMS": "0ms",
  "HistoryLink": "http://localhost:8080"
}
```

## Notes

If your target URL has HTTPS Janus will fail to send the webhook if the certificate is not valid.