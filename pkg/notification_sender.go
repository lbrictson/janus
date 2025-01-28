package pkg

import (
	"context"
	"errors"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/schema"
	"github.com/lbrictson/janus/pkg/notification_sender"
	"time"
)

func sendNotification(db *ent.Client, notificationChannelItem *ent.NotificationChannel, input notification_sender.NewNotificationInput) error {
	ctx := context.Background()
	switch notificationChannelItem.Type {
	case schema.ChannelSlack:
		sendErr := notification_sender.SendSlackNotification(input, notificationChannelItem.Config.SlackWebhook)
		if sendErr != nil {
			metricTrackFailedNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError(sendErr.Error()).
				SetLastUsed(time.Now()).
				Save(ctx)
		} else {
			metricTrackSentNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError("").
				SetLastUsed(time.Now()).
				Save(ctx)
		}
		return sendErr
	case schema.ChannelDiscord:
		sendErr := notification_sender.SendDiscordNotification(input, notificationChannelItem.Config.DiscordWebhookURL)
		if sendErr != nil {
			metricTrackFailedNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError(sendErr.Error()).
				SetLastUsed(time.Now()).
				Save(ctx)
		} else {
			metricTrackSentNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError("").
				SetLastUsed(time.Now()).
				Save(ctx)
		}
		return sendErr
	case schema.ChannelWebhook:
		sendErr := notification_sender.SendOutboundWebhookNotification(input, notificationChannelItem.Config.WebhookURL, notificationChannelItem.Config.WebhookHeaders)
		if sendErr != nil {
			metricTrackFailedNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError(sendErr.Error()).
				SetLastUsed(time.Now()).
				Save(ctx)
		} else {
			metricTrackSentNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError("").
				SetLastUsed(time.Now()).
				Save(ctx)
		}
		return sendErr
	case schema.ChannelAWSSNS:
		sendErr := notification_sender.SendSNSNotification(input, notificationChannelItem.Config.AWSCredentials.AccessKeyID, notificationChannelItem.Config.AWSCredentials.SecretAccessKey, notificationChannelItem.Config.AWSRegion, notificationChannelItem.Config.SNSTopicARN)
		if sendErr != nil {
			metricTrackFailedNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError(sendErr.Error()).
				SetLastUsed(time.Now()).
				Save(ctx)
		} else {
			metricTrackSentNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError("").
				SetLastUsed(time.Now()).
				Save(ctx)
		}
		return sendErr
	case schema.ChannelAWSEventBridge:
		sendErr := notification_sender.SendEventBridgeNotification(input, notificationChannelItem.Config.AWSCredentials.AccessKeyID, notificationChannelItem.Config.AWSCredentials.SecretAccessKey, notificationChannelItem.Config.AWSRegion, notificationChannelItem.Config.EventBusName, notificationChannelItem.Config.EventSource, notificationChannelItem.Config.DetailType)
		if sendErr != nil {
			metricTrackFailedNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError(sendErr.Error()).
				SetLastUsed(time.Now()).
				Save(ctx)
		} else {
			metricTrackSentNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError("").
				SetLastUsed(time.Now()).
				Save(ctx)
		}
		return sendErr
	case schema.ChannelEmail:
		c, _ := LoadConfig()
		s, _ := getSMTPConfig(ctx, db)
		sendErr := notification_sender.SendEmailNotification(
			notification_sender.NewNotificationInput{
				JobName:     input.JobName,
				ProjectName: input.ProjectName,
				JobStatus:   input.JobStatus,
				JobDuration: input.JobDuration,
				CallbackURL: input.CallbackURL,
			},
			*s,
			notificationChannelItem.Config.ToAddresses,
			c.ServerURL,
		)
		if sendErr != nil {
			metricTrackFailedNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError(sendErr.Error()).
				SetLastUsed(time.Now()).
				Save(ctx)
		} else {
			metricTrackSentNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError("").
				SetLastUsed(time.Now()).
				Save(ctx)
		}
		return sendErr
	case schema.ChannelTwilioSMS:
		sendErr := notification_sender.SendTwilioNotification(input, notificationChannelItem)
		if sendErr != nil {
			metricTrackFailedNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError(sendErr.Error()).
				SetLastUsed(time.Now()).
				Save(ctx)
		} else {
			metricTrackSentNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError("").
				SetLastUsed(time.Now()).
				Save(ctx)
		}
		return sendErr
	case schema.ChannelPagerDuty:
		sendErr := notification_sender.SendPagerdutyNotification(input, notificationChannelItem)
		if sendErr != nil {
			metricTrackFailedNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError(sendErr.Error()).
				SetLastUsed(time.Now()).
				Save(ctx)
		} else {
			metricTrackSentNotification()
			db.NotificationChannel.UpdateOne(notificationChannelItem).
				SetLastError("").
				SetLastUsed(time.Now()).
				Save(ctx)
		}
		return sendErr

	default:
		db.NotificationChannel.UpdateOne(notificationChannelItem).
			SetLastError("Unsupported notification channel type").
			SetLastUsed(time.Now()).
			Save(ctx)
		metricTrackFailedNotification()
		return errors.New("Unsupported notification channel type")
	}
}
