package notification_sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lbrictson/janus/ent"
	"net/http"
)

// PagerDutyEvent represents the payload for triggering a PagerDuty event
type PagerDutyEvent struct {
	RoutingKey  string       `json:"routing_key"`
	EventAction string       `json:"event_action"`
	Payload     EventPayload `json:"payload"`
}

// EventPayload represents the details of the incident
type EventPayload struct {
	Summary  string `json:"summary"`
	Source   string `json:"source"`
	Severity string `json:"severity"`
}

func SendPagerdutyNotification(input NewNotificationInput, config *ent.NotificationChannel) error {
	integrationKey := config.Config.PagerDutyToken
	// Define the event payload
	event := PagerDutyEvent{
		RoutingKey:  integrationKey,
		EventAction: "trigger",
		Payload: EventPayload{
			Summary:  fmt.Sprintf("%s in project %s : %s", input.JobName, input.ProjectName, input.JobStatus),
			Source:   "Janus",
			Severity: "critical",
		},
	}

	// Convert the event to JSON
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON for pagerduty notification: %v", err)
	}

	// Send the event to PagerDuty
	resp, err := http.Post(
		"https://events.pagerduty.com/v2/enqueue",
		"application/json",
		bytes.NewBuffer(eventJSON),
	)
	if err != nil {
		return fmt.Errorf("failed to send pagerduty notification: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unexpected status code from pagerduty api: %v", resp.StatusCode)
	}
	return nil
}
