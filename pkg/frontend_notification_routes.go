package pkg

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/schema"
	"github.com/lbrictson/janus/pkg/notification_sender"
)

func renderNotificationPage(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type FrontendChannel struct {
			ent.NotificationChannel
			LastUsed string
		}
		n, err := db.NotificationChannel.Query().All(c.Request().Context())
		if err != nil {
			slog.Error("error getting notification channels", "error", err)
			return renderErrorPage(c, "Error getting notification channels", http.StatusInternalServerError)
		}

		channels := make([]FrontendChannel, 0, len(n))
		for _, channel := range n {
			lastUsed := humanize.Time(channel.LastUsed)
			if channel.LastUsed.IsZero() {
				lastUsed = "Never"
			}
			channels = append(channels, FrontendChannel{
				NotificationChannel: *channel,
				LastUsed:            lastUsed,
			})
		}

		return c.Render(http.StatusOK, "notification-channels", map[string]any{
			"Channels": channels,
		})
	}
}

func renderNewNotificationPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "notification-channel-create", nil)
	}
}

func formCreateNotificationChannel(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Basic form structure
		type Form struct {
			Name        string `form:"name"`
			Description string `form:"description"`
			Type        string `form:"type"`
		}

		// Parse basic fields
		var form Form
		if err := c.Bind(&form); err != nil {
			return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
				"Error": "Invalid form data",
			})
		}
		// Begin building channel config based on type
		var channelConfig schema.ChannelConfig

		// Process type-specific configuration
		switch form.Type {
		case "discord":
			channelConfig.DiscordWebhookURL = c.FormValue("config.discord_webhook_url")
			if channelConfig.DiscordWebhookURL == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "Discord webhook URL is required",
				})
			}

		case "slack":
			channelConfig.SlackWebhook = c.FormValue("config.slack_webhook")
			if channelConfig.SlackWebhook == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "Slack webhook required",
				})
			}

		case "email":
			channelConfig.FromAddress = c.FormValue("config.from_address")
			toAddresses := c.FormValue("config.to_addresses")
			if channelConfig.FromAddress == "" || toAddresses == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "From address and to addresses are required",
				})
			}
			// Split to addresses by newline
			channelConfig.ToAddresses = strings.Split(strings.ReplaceAll(toAddresses, "\r\n", "\n"), "\n")

		case "teams":
			channelConfig.TeamsWebhookURL = c.FormValue("config.teams_webhook_url")
			if channelConfig.TeamsWebhookURL == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "Teams webhook URL is required",
				})
			}

		case "webhook":
			channelConfig.WebhookURL = c.FormValue("config.webhook_url")
			headers := c.FormValue("config.webhook_headers")
			if channelConfig.WebhookURL == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "Webhook URL is required",
				})
			}
			// Parse headers JSON if provided
			if headers != "" {
				if err := json.Unmarshal([]byte(headers), &channelConfig.WebhookHeaders); err != nil {
					return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
						"Error": "Invalid webhook headers JSON",
					})
				}
			}

		case "pagerduty":
			channelConfig.PagerDutyToken = c.FormValue("config.pagerduty_token")
			channelConfig.PagerDutyService = ""
			if channelConfig.PagerDutyToken == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "PagerDuty token and service are required",
				})
			}

		case "twilio-sms":
			channelConfig.TwilioAccountSID = c.FormValue("config.twilio_account_sid")
			channelConfig.TwilioAuthToken = c.FormValue("config.twilio_auth_token")
			channelConfig.TwilioFromNumber = c.FormValue("config.twilio_from_number")
			toNumbers := c.FormValue("config.twilio_to_numbers")
			if channelConfig.TwilioAccountSID == "" || channelConfig.TwilioAuthToken == "" ||
				channelConfig.TwilioFromNumber == "" || toNumbers == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "All Twilio fields are required",
				})
			}
			// Split to numbers by newline
			channelConfig.TwilioToNumbers = strings.Split(strings.ReplaceAll(toNumbers, "\r\n", "\n"), "\n")

		case "aws-sns":
			channelConfig.AWSRegion = c.FormValue("config.aws_region")
			channelConfig.AWSCredentials = schema.AWSCredentials{
				AccessKeyID:     c.FormValue("config.aws_credentials.access_key_id"),
				SecretAccessKey: c.FormValue("config.aws_credentials.secret_access_key"),
			}
			if channelConfig.AWSRegion == "" || channelConfig.AWSCredentials.AccessKeyID == "" ||
				channelConfig.AWSCredentials.SecretAccessKey == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "AWS region and credentials are required",
				})
			}
			channelConfig.SNSTopicARN = c.FormValue("config.sns_topic_arn")
			if channelConfig.SNSTopicARN == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "SNS topic ARN is required",
				})
			}

		case "aws-eventbridge":
			channelConfig.AWSRegion = c.FormValue("config.eventbridge_aws_region")
			channelConfig.AWSCredentials = schema.AWSCredentials{
				AccessKeyID:     c.FormValue("config.eventbridge_aws_credentials.access_key_id"),
				SecretAccessKey: c.FormValue("config.eventbridge_aws_credentials.secret_access_key"),
			}
			if channelConfig.AWSRegion == "" || channelConfig.AWSCredentials.AccessKeyID == "" ||
				channelConfig.AWSCredentials.SecretAccessKey == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "AWS region and credentials are required",
				})
			}
			channelConfig.EventBusName = c.FormValue("config.event_bus_name")
			channelConfig.EventSource = c.FormValue("config.event_source")
			channelConfig.DetailType = c.FormValue("config.detail_type")
			if channelConfig.EventBusName == "" || channelConfig.EventSource == "" || channelConfig.DetailType == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "All EventBridge fields are required",
				})
			}

		default:
			return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
				"Error": "Invalid notification channel type",
			})
		}

		// Create the notification channel
		channel, err := db.NotificationChannel.Create().
			SetName(form.Name).
			SetDescription(form.Description).
			SetType(schema.NotificationChannelType(form.Type)).
			SetConfig(channelConfig).
			Save(c.Request().Context())

		if err != nil {
			slog.Error("failed to create notification channel",
				"error", err,
				"name", form.Name,
				"type", form.Type,
			)
			return c.Render(http.StatusInternalServerError, "notification-channel-create", map[string]interface{}{
				"Error": "Failed to create notification channel",
			})
		}

		slog.Info("created notification channel",
			"id", channel.ID,
			"name", channel.Name,
			"type", channel.Type,
		)

		return c.Redirect(http.StatusSeeOther, "/notifications")
	}
}

func renderNotificationChannelEditPage(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		channelID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return renderErrorPage(c, "Invalid channel ID", http.StatusBadRequest)
		}

		channel, err := db.NotificationChannel.Get(c.Request().Context(), channelID)
		if err != nil {
			slog.Error("error getting notification channel", "error", err)
			return renderErrorPage(c, "Error getting notification channel", http.StatusInternalServerError)
		}

		return c.Render(http.StatusOK, "edit-notification-channel", map[string]any{
			"Channel": channel,
		})
	}
}

func hookNotificationToggleStatus(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		channelID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			slog.Error("error parsing channel ID", "error", err)
			return renderErrorPage(c, "Invalid channel ID", http.StatusBadRequest)
		}

		channel, err := db.NotificationChannel.Get(c.Request().Context(), channelID)
		if err != nil {
			slog.Error("error getting notification channel", "error", err)
			return renderErrorPage(c, "Error getting notification channel", http.StatusInternalServerError)
		}
		if channel.Enabled {
			channel.Update().SetEnabled(false).Save(c.Request().Context())
		} else {
			channel.Update().SetEnabled(true).Save(c.Request().Context())
		}
		return c.JSON(http.StatusOK, map[string]any{"status": "ok"})
	}
}

func deleteNotificationChannel(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		channelID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			slog.Error("error parsing channel ID", "error", err)
			return renderErrorPage(c, "Invalid channel ID", http.StatusBadRequest)
		}

		err = db.NotificationChannel.DeleteOneID(channelID).Exec(c.Request().Context())
		if err != nil {
			slog.Error("error deleting notification channel", "error", err)
			return renderErrorPage(c, "Error deleting notification channel", http.StatusInternalServerError)
		}
		// Now delete it for any job that has it
		jobsWithThisChannel, err := db.Job.Query().All(c.Request().Context())
		if err != nil {
			slog.Error("error getting jobs", "error", err)
			return renderErrorPage(c, "Error getting jobs", http.StatusInternalServerError)
		}
		for _, j := range jobsWithThisChannel {
			needsUpdate := false
			for i, channel := range j.NotifyOnFailureChannelIds {
				if channel == channelID {
					// Remove offending channel from array
					j.NotifyOnFailureChannelIds = append(j.NotifyOnFailureChannelIds[:i], j.NotifyOnFailureChannelIds[i+1:]...)
					needsUpdate = true
				}
			}
			for i, channel := range j.NotifyOnSuccessChannelIds {
				if channel == channelID {
					// Remove offending channel from array
					j.NotifyOnSuccessChannelIds = append(j.NotifyOnSuccessChannelIds[:i], j.NotifyOnSuccessChannelIds[i+1:]...)
					needsUpdate = true
				}
			}
			for i, channel := range j.NotifyOnFailureChannelIds {
				if channel == channelID {
					// Remove offending channel from array
					j.NotifyOnFailureChannelIds = append(j.NotifyOnFailureChannelIds[:i], j.NotifyOnFailureChannelIds[i+1:]...)
					needsUpdate = true
				}
			}
			if needsUpdate {
				db.Job.UpdateOne(j).SetNotifyOnFailureChannelIds(j.NotifyOnFailureChannelIds).SetNotifyOnSuccessChannelIds(j.NotifyOnSuccessChannelIds).SetNotifyOnFailureChannelIds(j.NotifyOnFailureChannelIds).Exec(c.Request().Context())
			}
		}
		return c.Redirect(http.StatusSeeOther, "/notifications")
	}
}

func formEditNotificationChannel(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		channelID := c.Param("id")
		ctx := c.Request().Context()
		channelIDInt, err := strconv.Atoi(channelID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid channel ID")
		}
		// Get existing channel
		channel, err := db.NotificationChannel.Get(ctx, channelIDInt)
		if err != nil {
			if ent.IsNotFound(err) {
				return echo.NewHTTPError(http.StatusNotFound, "Channel not found")
			}
			return err
		}

		// For GET requests, display the form
		if c.Request().Method == "GET" {
			return c.Render(http.StatusOK, "edit-notification-channel", map[string]interface{}{
				"Channel": channel,
			})
		}

		// Basic form structure
		type Form struct {
			Name        string `form:"name"`
			Description string `form:"description"`
			Type        string `form:"type"`
		}

		// Parse basic fields
		var form Form
		if err := c.Bind(&form); err != nil {
			return c.Render(http.StatusBadRequest, "edit-notification-channel", map[string]interface{}{
				"Channel": channel,
				"Error":   "Invalid form data",
			})
		}
		// Begin building updated channel config
		var channelConfig schema.ChannelConfig

		// Process type-specific configuration
		switch form.Type {
		case "discord":
			channelConfig.DiscordWebhookURL = c.FormValue("config.discord_webhook_url")
			if channelConfig.DiscordWebhookURL == "" {
				return c.Render(http.StatusBadRequest, "edit-notification-channel", map[string]interface{}{
					"Channel": channel,
					"Error":   "Discord webhook URL is required",
				})
			}

		case "slack":
			channelConfig.SlackWebhook = c.FormValue("config.slack_webhook")
			if channelConfig.SlackWebhook == "" {
				return c.Render(http.StatusBadRequest, "edit-notification-channel", map[string]interface{}{
					"Channel": channel,
					"Error":   "Slack webhook is required",
				})
			}

		case "email":
			channelConfig.FromAddress = c.FormValue("config.from_address")
			toAddresses := c.FormValue("config.to_addresses")
			if channelConfig.FromAddress == "" || toAddresses == "" {
				return c.Render(http.StatusBadRequest, "edit-notification-channel", map[string]interface{}{
					"Channel": channel,
					"Error":   "From address and to addresses are required",
				})
			}
			// Split to addresses by newline
			channelConfig.ToAddresses = strings.Split(strings.ReplaceAll(toAddresses, "\r\n", "\n"), "\n")

		case "teams":
			channelConfig.TeamsWebhookURL = c.FormValue("config.teams_webhook_url")
			if channelConfig.TeamsWebhookURL == "" {
				return c.Render(http.StatusBadRequest, "edit-notification-channel", map[string]interface{}{
					"Channel": channel,
					"Error":   "Teams webhook URL is required",
				})
			}

		case "webhook":
			channelConfig.WebhookURL = c.FormValue("config.webhook_url")
			headers := c.FormValue("config.webhook_headers")
			if channelConfig.WebhookURL == "" {
				return c.Render(http.StatusBadRequest, "edit-notification-channel", map[string]interface{}{
					"Channel": channel,
					"Error":   "Webhook URL is required",
				})
			}
			// Parse headers JSON if provided
			if headers != "" {
				if err := json.Unmarshal([]byte(headers), &channelConfig.WebhookHeaders); err != nil {
					return c.Render(http.StatusBadRequest, "edit-notification-channel", map[string]interface{}{
						"Channel": channel,
						"Error":   "Invalid webhook headers JSON",
					})
				}
			}

		case "pagerduty":
			channelConfig.PagerDutyToken = c.FormValue("config.pagerduty_token")
			if channelConfig.PagerDutyToken == "" {
				return c.Render(http.StatusBadRequest, "edit-notification-channel", map[string]interface{}{
					"Channel": channel,
					"Error":   "PagerDuty token and service are required",
				})
			}

		case "twilio-sms":
			channelConfig.TwilioAccountSID = c.FormValue("config.twilio_account_sid")
			channelConfig.TwilioAuthToken = c.FormValue("config.twilio_auth_token")
			channelConfig.TwilioFromNumber = c.FormValue("config.twilio_from_number")
			toNumbers := c.FormValue("config.twilio_to_numbers")
			if channelConfig.TwilioAccountSID == "" ||
				channelConfig.TwilioFromNumber == "" || toNumbers == "" {
				return c.Render(http.StatusBadRequest, "edit-notification-channel", map[string]interface{}{
					"Channel": channel,
					"Error":   "All Twilio fields except auth token are required",
				})
			}
			// If auth token is empty, keep the existing one
			if channelConfig.TwilioAuthToken == "" {
				channelConfig.TwilioAuthToken = channel.Config.TwilioAuthToken
			}
			// Split to numbers by newline
			channelConfig.TwilioToNumbers = strings.Split(strings.ReplaceAll(toNumbers, "\r\n", "\n"), "\n")

		case "aws-sns":
			channelConfig.AWSRegion = c.FormValue("config.aws_region")
			channelConfig.AWSCredentials = schema.AWSCredentials{
				AccessKeyID:     c.FormValue("config.aws_credentials.access_key_id"),
				SecretAccessKey: c.FormValue("config.aws_credentials.secret_access_key"),
			}
			if channelConfig.AWSRegion == "" || channelConfig.AWSCredentials.AccessKeyID == "" ||
				channelConfig.AWSCredentials.SecretAccessKey == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "AWS region and credentials are required",
				})
			}
			channelConfig.SNSTopicARN = c.FormValue("config.sns_topic_arn")
			if channelConfig.SNSTopicARN == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "SNS topic ARN is required",
				})
			}

		case "aws-eventbridge":
			channelConfig.AWSRegion = c.FormValue("config.eventbridge_aws_region")
			channelConfig.AWSCredentials = schema.AWSCredentials{
				AccessKeyID:     c.FormValue("config.eventbridge_aws_credentials.access_key_id"),
				SecretAccessKey: c.FormValue("config.eventbridge_aws_credentials.secret_access_key"),
			}
			if channelConfig.AWSRegion == "" || channelConfig.AWSCredentials.AccessKeyID == "" ||
				channelConfig.AWSCredentials.SecretAccessKey == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "AWS region and credentials are required",
				})
			}
			channelConfig.EventBusName = c.FormValue("config.event_bus_name")
			channelConfig.EventSource = c.FormValue("config.event_source")
			channelConfig.DetailType = c.FormValue("config.detail_type")
			if channelConfig.EventBusName == "" || channelConfig.EventSource == "" || channelConfig.DetailType == "" {
				return c.Render(http.StatusBadRequest, "notification-channel-create", map[string]interface{}{
					"Error": "All EventBridge fields are required",
				})
			}

		default:
			return c.Render(http.StatusBadRequest, "edit-notification-channel", map[string]interface{}{
				"Channel": channel,
				"Error":   "Invalid notification channel type",
			})
		}

		// Update the notification channel
		updatedChannel, err := channel.Update().
			SetName(form.Name).
			SetDescription(form.Description).
			SetConfig(channelConfig).
			Save(ctx)

		if err != nil {
			slog.Error("failed to update notification channel",
				"error", err,
				"channel_id", channel.ID,
				"name", form.Name,
			)
			return c.Render(http.StatusInternalServerError, "edit-notification-channel", map[string]interface{}{
				"Channel": channel,
				"Error":   "Failed to update notification channel",
			})
		}

		slog.Info("updated notification channel",
			"channel_id", updatedChannel.ID,
			"name", updatedChannel.Name,
			"type", updatedChannel.Type,
		)

		return c.Redirect(http.StatusSeeOther, "/notifications")
	}
}

func hookSendTestNotification(db *ent.Client, config Config) echo.HandlerFunc {
	return func(c echo.Context) error {

		channelID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			slog.Error("error parsing channel ID", "error", err)
			return c.String(http.StatusBadRequest, "invalid channel ID")
		}

		channel, err := db.NotificationChannel.Get(c.Request().Context(), channelID)
		if err != nil {
			slog.Error("error getting notification channel", "error", err)
			return c.String(http.StatusInternalServerError, "error getting notification channel")
		}

		// Send a test notification
		err = sendNotification(db, channel, notification_sender.NewNotificationInput{
			JobName:     "TEST",
			ProjectName: "TEST",
			JobStatus:   notification_sender.TESTING,
			JobDuration: "0ms",
			CallbackURL: config.ServerURL,
		})
		if err != nil {
			slog.Error("error sending test notification", "error", err)
			return c.String(http.StatusInternalServerError, "error sending test notification")
		}

		return c.JSON(http.StatusOK, map[string]any{"status": "ok"})
	}
}
