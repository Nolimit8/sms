package job

import (
	"net/http"
	"os"
)

func RunJob(w http.ResponseWriter, _ *http.Request) {
	newPostApiKey := os.Getenv("NP_API_KEY")
	newPostService := NewPostServiceImpl{
		apiKey: newPostApiKey,
	}
	notificationsService := SMSClubSMSNotificationsService{
		username: os.Getenv("SMS_CLUB_USERNAME"),
		password: os.Getenv("SMS_CLUB_PASSWORD"),
	}
	job := ReminderJob{
		newPostService:           newPostService,
		smsNotificationsService:  notificationsService,
		notificationTextProvider: DaysAwareMessageProvider{},
	}
	jobError := job.RunJob()
	for i := range jobError {
		w.Write([]byte(jobError[i].Error()))
	}
	w.Write([]byte("Job completed successfully"))
}
