package notification_sender

type NotificationStatus string

const (
	SUCCESS  = NotificationStatus("Success")
	FAILURE  = NotificationStatus("Failure")
	STARTING = NotificationStatus("Starting")
	TESTING  = NotificationStatus("Testing")
)

type NewNotificationInput struct {
	JobName     string
	ProjectName string
	JobStatus   NotificationStatus
	JobDuration string
	CallbackURL string
}
