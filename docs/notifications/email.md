# Email

Janus sends email notifications using the SMTP protocol.  To enable email notifications, you will need to configure Janus with your SMTP server settings in the admin configuration.  
Changing the SMTP settings requires the global admin permission.

Create a new PagerDuty channel at `/notifications`

## Configuration

| Config Item  | Explanation                                                                   |
|--------------|-------------------------------------------------------------------------------|
| From Address | The address to place in the from field                                        |
| To Addresses | One or more email addresses to send notifications to.  A new line per address |

## Notes

Ensure your SMTP server is configured for DKIM and SPF to prevent emails from being marked as spam.