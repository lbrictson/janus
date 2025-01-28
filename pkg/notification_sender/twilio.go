package notification_sender

import (
	"fmt"
	"github.com/lbrictson/janus/ent"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendTwilioNotification(input NewNotificationInput, config *ent.NotificationChannel) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.Config.TwilioAccountSID,
		Password: config.Config.TwilioAuthToken,
	})
	content := fmt.Sprintf("%s in project %s : %s\nFull Output: %v\nStatus: %v", input.JobName, input.ProjectName, input.JobStatus, input.CallbackURL, input.JobStatus)
	for _, x := range config.Config.TwilioToNumbers {
		params := &openapi.CreateMessageParams{}
		params.SetTo(x)
		params.SetFrom(config.Config.TwilioFromNumber)
		params.SetBody(content)

		_, err := client.Api.CreateMessage(params)
		if err != nil {
			return fmt.Errorf("failed to send twilio notification: %v", err)
		}
	}
	return nil
}
