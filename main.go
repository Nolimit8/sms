package main

import "os"

func main() {
	/*	newPostApiKey := os.Getenv("NP_API_KEY")
		newPostService := NewPostServiceImpl{
			apiKey: newPostApiKey,
		}
		resp, err := newPostService.GetAllInternetDocuments()
		if err != nil {
			print(resp)
		} else {
			print(err)
		}*/

	notificationsService := SMSClubSMSNotificationsService{
		username: os.Getenv("SMS_CLUB_USERNAME"),
		password: os.Getenv("SMS_CLUB_PASSWORD"),
	}
	err := notificationsService.SendSMSBatch(SMSBatchNotification{
		PhoneNumbers: []string{"380637240604"},
		Message:      "Hello world",
	})
	print(err)
}
