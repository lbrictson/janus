package notification_sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DiscordPayload struct {
	Content string `json:"content"`
}

func SendDiscordNotification(input NewNotificationInput, webhookURL string) error {
	payload := DiscordPayload{
		Content: fmt.Sprintf("%s in project %s : %s\nFull Output: %v\nStatus: %v", input.JobName, input.ProjectName, input.JobStatus, input.CallbackURL, input.JobStatus),
	}
	if input.JobStatus != STARTING {
		payload.Content += fmt.Sprintf("\nDuration: %v", input.JobDuration)
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	return nil
}
