package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
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
	SendSMSBatch(message SMSMessage) error
}

type SMSClubSMSNotificationsService struct {
	username string
	password string
}

func (service SMSClubSMSNotificationsService) SendSMSBatch(notification SMSBatchNotification) error {
	apiEndpoint := "https://sms-fly.ua/api/api.php"
	client := &http.Client{}
	requestBody := SMSRequest{
		Operation: "SENDSMS",
		Message: SMSMessage{
			StartTime:   "AUTO",
			EndTime:     "AUTO",
			Lifetime:    "4",
			Rate:        "120",
			Description: "Daily notification job",
			Source:      "putivoditel",
			Body:        notification.Message,
			Recipient:   notification.PhoneNumbers,
		},
	}
	requestBodyBytes, marshalError := xml.MarshalIndent(requestBody, "", "")
	if marshalError != nil {
		return marshalError
	}
	print(string(requestBodyBytes))
	request, requestError := http.NewRequest("POST", apiEndpoint, bytes.NewReader([]byte(xml.Header+string(requestBodyBytes))))
	if requestError != nil {
		return requestError
	}
	request.SetBasicAuth(service.username, service.password)
	resp, err := client.Do(request)
	if resp != nil {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		print(string(bodyBytes))
	}
	return err
}
