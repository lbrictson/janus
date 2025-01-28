package notification_sender

type NotificationStatus string

const (
	SUCCESS  = NotificationStatus("Success")
	FAILURE  = NotificationStatus("Failure")
	STARTING = NotificationStatus("Starting")
	TESTING  = NotificationStatus("Testing")
)

type NewNotificationInput struct {
	JobName     string             `json:"job_name"`
	ProjectName string             `json:"project_name"`
	JobStatus   NotificationStatus `json:"job_status"`
	JobDuration string             `json:"job_duration"`
	CallbackURL string             `json:"callback_url"`
}
