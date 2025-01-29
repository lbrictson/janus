# PagerDuty

Janus can integrate with PagerDuty to send notifications when jobs fail.  You create a PagerDuty notification channel per service you wish to trigger in Pagerduty

Create a new PagerDuty channel at `/notifications`

## Configuration

| Config Item     | Explanation                                                 |
|-----------------|-------------------------------------------------------------|
| Integration Key | The integration key from PagerDuty for a particular service |

## Notes

You can learn more about PagerDuty integration keys [here](https://support.pagerduty.com/docs/services-and-integrations#section-integration-keys).

You must have the Events API V2 integration enabled for your service to use this feature.
