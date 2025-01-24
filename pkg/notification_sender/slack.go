package notification_sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SlackPayload struct {
	Blocks []SlackBlock `json:"blocks"`
}

type SlackBlock struct {
	Type string       `json:"type"`
	Text SlackSection `json:"text,omitempty"`
}

type SlackSection struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

func SendSlackNotification(input NewNotificationInput, webhookURL string) error {
	p := []SlackBlock{
		{
			Type: "header",
			Text: SlackSection{
				Type: "plain_text",
				Text: fmt.Sprintf("%s in project %s : %s", input.JobName, input.ProjectName, input.JobStatus),
			},
		},
		{
			Type: "section",
			Text: SlackSection{
				Type: "mrkdwn",
				Text: fmt.Sprintf("<%v|Full Output>", input.CallbackURL),
			},
		},
		{
			Type: "section",
			Text: SlackSection{
				Type: "mrkdwn",
				Text: fmt.Sprintf("Status: %v", input.JobStatus),
			},
		},
	}
	if input.JobStatus != STARTING {
		durationBlock := SlackBlock{
			Type: "section",
			Text: SlackSection{
				Type: "mrkdwn",
				Text: fmt.Sprintf("Duration: %v", input.JobDuration),
			},
		}
		p = append(p, durationBlock)
	}
	payload := SlackPayload{
		Blocks: p,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Send the payload to Slack
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error sending slack notification: %s", resp.Status)
	}
	return nil
}
