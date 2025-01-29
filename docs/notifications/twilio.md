# Twilio (SMS)

Twilio is a service that allows you to send and receive SMS messages. It is a paid service that is generally pay as you go.
You should budget for the cost of sending SMS messages when using this service.

Create a new Slack channel at `/notifications`

## Configuration

| Config Item | Explanation |
|-------------|-------------|
| Account SID | The Twilio Account SID |
| Auth Token  | The Twilio Auth Token |
| From Number | The Twilio phone number to send messages from |
| To Numbers  | One or more phone numbers to send messages to.  A new line per number |

## Notes

If you are in the United States you need to fully register your Twilio phone number to be able to send SMS messages, otherwise 
they will fail.

You can learn more about Twilio [here](https://www.twilio.com/docs/sms/send-messages).
```