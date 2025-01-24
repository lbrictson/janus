package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// NotificationChannelType represents the type of notification channel
type NotificationChannelType string

const (
	ChannelDiscord        NotificationChannelType = "discord"
	ChannelSlack          NotificationChannelType = "slack"
	ChannelEmail          NotificationChannelType = "email"
	ChannelTeams          NotificationChannelType = "teams"
	ChannelWebhook        NotificationChannelType = "webhook"
	ChannelPagerDuty      NotificationChannelType = "pagerduty"
	ChannelTwilioSMS      NotificationChannelType = "twilio_sms"
	ChannelAWSSNS         NotificationChannelType = "aws_sns"
	ChannelAWSEventBridge NotificationChannelType = "aws_eventbridge"
)

// ChannelConfig holds the type-specific configuration
type ChannelConfig struct {
	// Discord
	DiscordWebhookURL string `json:"webhook_url,omitempty"`

	// Slack
	SlackWebhook string `json:"slack_webhook,omitempty"`

	// Email
	FromAddress string   `json:"from_address,omitempty"`
	ToAddresses []string `json:"to_addresses,omitempty"`

	// Teams
	TeamsWebhookURL string `json:"teams_webhook_url,omitempty"`

	// Webhook
	WebhookURL     string            `json:"url,omitempty"`
	WebhookHeaders map[string]string `json:"headers,omitempty"`

	// PagerDuty
	PagerDutyToken   string `json:"pagerduty_token,omitempty"`
	PagerDutyService string `json:"pagerduty_service,omitempty"`

	// Twilio
	TwilioAccountSID string   `json:"twilio_account_sid,omitempty"`
	TwilioAuthToken  string   `json:"twilio_auth_token,omitempty"`
	TwilioFromNumber string   `json:"twilio_from_number,omitempty"`
	TwilioToNumbers  []string `json:"twilio_to_numbers,omitempty"`

	// AWS SNS
	SNSTopicARN    string         `json:"sns_topic_arn,omitempty"`
	AWSRegion      string         `json:"aws_region,omitempty"`
	AWSCredentials AWSCredentials `json:"aws_credentials,omitempty"`

	// AWS EventBridge
	EventBusName string `json:"event_bus_name,omitempty"`
	EventSource  string `json:"event_source,omitempty"`
	DetailType   string `json:"detail_type,omitempty"`
}

type AWSCredentials struct {
	AccessKeyID     string `json:"access_key_id,omitempty"`
	SecretAccessKey string `json:"secret_access_key,omitempty"`
	RoleARN         string `json:"role_arn,omitempty"`
}

// NotificationChannel holds the schema definition for the NotificationChannel entity.
type NotificationChannel struct {
	ent.Schema
}

// Fields of the NotificationChannel.
func (NotificationChannel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("description").
			Optional(),
		field.Enum("type").
			GoType(NotificationChannelType("")),
		field.JSON("config", ChannelConfig{}).
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Bool("enabled").
			Default(true),
		field.Int("retry_count").
			Optional().
			Default(3),
		field.Time("last_used").
			Optional(),
		field.String("last_error").
			Optional(),
	}
}

// Values provides the list of valid notification channel types
func (NotificationChannelType) Values() []string {
	return []string{
		string(ChannelDiscord),
		string(ChannelSlack),
		string(ChannelEmail),
		string(ChannelTeams),
		string(ChannelWebhook),
		string(ChannelPagerDuty),
		string(ChannelTwilioSMS),
		string(ChannelAWSSNS),
		string(ChannelAWSEventBridge),
	}
}
