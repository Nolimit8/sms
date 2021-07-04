package job

import (
	"bytes"
	"errors"
	"text/template"
	"time"
)

type MessageProvider interface {
	GetMessageForInternetDocumentsAwaitingPickup(document InternetDocument) (string, error)
	GetMessageForCreatedInternetDocument(document InternetDocument) (string, error)
}

type DaysAwareMessageProvider struct {
}

func (provider DaysAwareMessageProvider) GetMessageForInternetDocumentsAwaitingPickup(document InternetDocument) (string, error) {
	statusUpdateDate, _ := time.Parse("2006-01-02 15:04:05", document.StatusUpdateDate)
	currentDateTime := time.Now()
	timeFromLastStatusUpdate := int(currentDateTime.Sub(statusUpdateDate).Hours() / 24)
	if timeFromLastStatusUpdate >= 0 && timeFromLastStatusUpdate <= 8 {
		t, _ := template.New("smsTemplate").Parse(DocumentsAwaitingPickupSMSTemplate(timeFromLastStatusUpdate).GetTemplate())
		var tpl bytes.Buffer
		if err := t.Execute(&tpl, document); err != nil {
			return "", err
		}
		return tpl.String(), nil
	} else {
		return "", errors.New("invalid date")
	}
}

func (provider DaysAwareMessageProvider) GetMessageForCreatedInternetDocument(document InternetDocument) (string, error) {
	statusUpdateDate, _ := time.Parse("2006-01-02 15:04:05", document.StatusUpdateDate)
	currentDateTime := time.Now()
	timeFromLastStatusUpdate := currentDateTime.Sub(statusUpdateDate).Hours() / 24
	if timeFromLastStatusUpdate <= 1 {
		t, _ := template.New("smsTemplate").Parse(DispatchedDocumentSMSTemplate)
		var tpl bytes.Buffer
		if err := t.Execute(&tpl, document); err != nil {
			return "", err
		}
		return tpl.String(), nil
	} else {
		return "", errors.New("invalid date")
	}
}
