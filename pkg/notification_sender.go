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

	default:
		db.NotificationChannel.UpdateOne(notificationChannelItem).
			SetLastError("Unsupported notification channel type").
			SetLastUsed(time.Now()).
			Save(ctx)
		metricTrackFailedNotification()
		return errors.New("Unsupported notification channel type")
	}
}
