package pkg

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	notificationsSent = promauto.NewCounter(prometheus.CounterOpts{
		Name: "janus_notifications_sent_total",
		Help: "The total number of notifications sent",
	})
	notificationsFailure = promauto.NewCounter(prometheus.CounterOpts{
		Name: "janus_notifications_failed_total",
		Help: "The total number of notifications that failed to send",
	})
	loginFailures = promauto.NewCounter(prometheus.CounterOpts{
		Name: "janus_login_failures_total",
		Help: "The total number of login failures",
	})
	loginSuccesses = promauto.NewCounter(prometheus.CounterOpts{
		Name: "janus_login_successes_total",
		Help: "The total number of successful logins",
	})
	totalJobRuns = promauto.NewCounter(prometheus.CounterOpts{
		Name: "janus_total_job_runs",
		Help: "The total number of job runs",
	})
	totalJobFailures = promauto.NewCounter(prometheus.CounterOpts{
		Name: "janus_total_job_failures",
		Help: "The total number of job failures",
	})
	totalJobSuccesses = promauto.NewCounter(prometheus.CounterOpts{
		Name: "janus_total_job_successes",
		Help: "The total number of job successes",
	})
)

func RunMetricsListener(port int) {
	slog.Info("starting metrics server", "port", port)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func metricTrackSentNotification() {
	notificationsSent.Inc()
}

func metricTrackFailedNotification() {
	notificationsFailure.Inc()
}
