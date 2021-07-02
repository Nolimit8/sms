package job

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

const (
	SmsClubApiAddress       = "https://sms-fly.ua/api/api.php"
	SendSmsOperationName    = "SENDSMS"
	DefaultSmsStartTime     = "AUTO"
	DefaultSmsEndTime       = "AUTO"
	DefaultSmsSenderName    = "putivoditel"
	DefaultSmsLifetime      = "4"
	DefaultSmsRate          = "120"
	SmsJobDescriptionFormat = "Daily notification job %s"
)

type SMSBatchNotification struct {
	PhoneNumbers []string
	Message      string
}

type SMSRequest struct {
	XMLName   xml.Name   `xml:"request"`
	Operation string     `xml:"operation"`
	Message   SMSMessage `xml:"message"`
}

type SMSMessage struct {
	StartTime   string   `xml:"start_time,attr"`
	EndTime     string   `xml:"end_time,attr"`
	Lifetime    string   `xml:"lifetime,attr"`
	Rate        string   `xml:"rate,attr"`
	Description string   `xml:"desc,attr"`
	Source      string   `xml:"source,attr"`
	Body        string   `xml:"body"`
	Recipient   []string `xml:"recipient"`
}

type SMSNotificationsService interface {
	SendSMSBatch(message SMSBatchNotification) error
}

type SMSClubSMSNotificationsService struct {
	username string
	password string
}

func (service SMSClubSMSNotificationsService) SendSMSBatch(notification SMSBatchNotification) error {
	client := &http.Client{}
	requestBody := SMSRequest{
		Operation: SendSmsOperationName,
		Message: SMSMessage{
			StartTime:   DefaultSmsStartTime,
			EndTime:     DefaultSmsEndTime,
			Lifetime:    DefaultSmsLifetime,
			Rate:        DefaultSmsRate,
			Description: fmt.Sprintf(SmsJobDescriptionFormat, time.Now().String()),
			Source:      DefaultSmsSenderName,
			Body:        notification.Message,
			Recipient:   notification.PhoneNumbers,
		},
	}
	requestBodyBytes, marshalError := xml.Marshal(requestBody)
	if marshalError != nil {
		return marshalError
	}
	print(string(requestBodyBytes))
	request, requestError := http.NewRequest(http.MethodPost, SmsClubApiAddress, bytes.NewReader([]byte(xml.Header+string(requestBodyBytes))))
	if requestError != nil {
		return requestError
	}
	request.SetBasicAuth(service.username, service.password)
	_, err := client.Do(request)
	return err
}
