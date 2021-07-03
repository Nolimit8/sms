package job

type Job interface {
	RunJob() error
}

type ReminderJob struct {
	newPostService           NewPostService
	smsNotificationsService  SMSNotificationsService
	notificationTextProvider MessageProvider
}

type InternetDocumentHandler interface {
	HandleInternetDocument(document InternetDocument) error
}

type InternetDocumentSMSDispatcher struct {
	smsNotificationsService  SMSNotificationsService
	notificationTextProvider MessageProvider
}

func (job ReminderJob) RunJob() []error {
	internetDocuments, internetDocumentsFetchingError := job.newPostService.GetAllInternetDocuments()
	if internetDocumentsFetchingError != nil {
		return []error{internetDocumentsFetchingError}
	}
	var errors []error
	internetDocumentsAwaitingPickup := job.newPostService.FilterInternetDocumentsAwaitingPickup(internetDocuments)
	for i := range internetDocumentsAwaitingPickup {
		currentInternetDocument := internetDocumentsAwaitingPickup[i]
		messageBody, messageGenerationError := job.notificationTextProvider.GetMessageForDispatchedInternetDocument(currentInternetDocument)
		if messageGenerationError == nil {
			smsDeliveryError := job.smsNotificationsService.SendSMSBatch(SMSBatchNotification{
				PhoneNumbers: []string{currentInternetDocument.RecipientContactPhone},
				Message:      messageBody,
			})
			if smsDeliveryError != nil {
				errors = append(errors, smsDeliveryError)
			}
		} else {
			errors = append(errors, messageGenerationError)
		}
	}
	return errors

}
