package notification_sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OutboundWebhookPayload struct {
	JobName       string
	JobID         int
	ProjectName   string
	JobStatus     NotificationStatus
	JobDurationMS string
	HistoryLink   string
}

func SendOutboundWebhookNotification(input NewNotificationInput, webhookURL string, headers map[string]string) error {
	payload := OutboundWebhookPayload{
		JobName:       input.JobName,
		ProjectName:   input.ProjectName,
		JobStatus:     input.JobStatus,
		JobDurationMS: input.JobDuration,
		HistoryLink:   input.CallbackURL,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 && resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code when sending outbound webhook: %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	return nil
}
