package job

type Job interface {
	RunJob() error
}

type ReminderJob struct {
	newPostService          NewPostService
	smsNotificationsService SMSNotificationsService
	notificationText        string
}

func (job ReminderJob) RunJob() error {
	internetDocuments, internetDocumentsFetchingError := job.newPostService.GetAllInternetDocuments()
	if internetDocumentsFetchingError != nil {
		return internetDocumentsFetchingError
	}
	internetDocumentsAwaitingPickup := job.newPostService.FilterInternetDocumentsAwaitingPickup(internetDocuments)
	var phoneNumbers []string
	for i := range internetDocumentsAwaitingPickup {
		phoneNumbers = append(phoneNumbers, internetDocuments[i].RecipientContactPhone)
	}
	smsDeliveryError := job.smsNotificationsService.SendSMSBatch(SMSBatchNotification{
		PhoneNumbers: phoneNumbers,
		Message:      job.notificationText,
	})
	return smsDeliveryError
}
