package notification_sender

import (
	"crypto/tls"
	"fmt"
	"github.com/lbrictson/janus/ent"
	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

func SendEmailNotification(input NewNotificationInput, config ent.SMTPConfig, toAddresses []string, serverURL string, brand string) error {
	// Configure hermes by setting a theme and your product info
	status := fmt.Sprintf("%v", input.JobStatus)
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: brand,
			Link: serverURL,
		},
	}
	tableData := [][]hermes.Entry{
		{
			// Key is the column name, Value is the cell value
			// First object defines what columns will be displayed
			{Key: "Job", Value: input.JobName},
			{Key: "Project", Value: input.ProjectName},
			{Key: "Status", Value: status},
		},
	}
	if input.JobStatus != STARTING {
		tableData = append(tableData, []hermes.Entry{
			{Key: "Duration", Value: input.JobDuration},
		})
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name:     input.JobName,
			Intros:   nil,
			Greeting: "Job status update:",
			Table: hermes.Table{
				Data: tableData,
			},
			Actions: []hermes.Action{
				{
					Button: hermes.Button{
						Color: "#38bdf8", // Optional action button color
						Text:  "View Job Details",
						Link:  input.CallbackURL,
					},
				},
			},
			Signature: "Automated Email Status",
			Outros:    []string{},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return err
	}
	d := gomail.NewDialer(config.SMTPServer, config.SMTPPort, config.SMTPUsername, config.SMTPPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	m := gomail.NewMessage()
	m.SetHeader("From", config.SMTPSender)
	m.SetHeader("To", toAddresses...)
	m.SetHeader("Subject", fmt.Sprintf("%v : %v in %v", input.JobStatus, input.JobStatus, input.ProjectName))
	m.SetBody("text/html", emailBody)
	err = d.DialAndSend(m)
	return err
}
