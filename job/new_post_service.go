package job

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type InternetDocument struct {
	ReferenceId           string `json:"Ref"`
	RecipientContactPhone string `json:"RecipientContactPhone"`
	DeliveryStatus        string `json:"StateId"`
	StatusUpdateDate      string `json:"DateLastUpdatedStatus"`
	RecipientName         string `json:"RecipientContactPerson"`
}

type NewPostRequest struct {
	ModelName        string            `json:"modelName"`
	CalledMethod     string            `json:"calledMethod"`
	MethodProperties map[string]string `json:"methodProperties"`
	ApiKey           string            `json:"apiKey"`
}

type NewPostService interface {
	GetAllInternetDocuments() ([]InternetDocument, error)
	FilterInternetDocumentsAwaitingPickup([]InternetDocument) []InternetDocument
	FilterDispatchedInternetDocuments([]InternetDocument) []InternetDocument
}

type GetInternetDocumentsResponse struct {
	Success bool               `json:"success"`
	Data    []InternetDocument `json:"data"`
	Errors  []string           `json:"errors"`
}

type NewPostServiceImpl struct {
	apiKey string
}

func (n NewPostServiceImpl) GetAllInternetDocuments() ([]InternetDocument, error) {
	requestUrl := "https://api.novaposhta.ua/v2.0/json/"
	requestBody := NewPostRequest{
		ModelName:    "InternetDocument",
		CalledMethod: "getDocumentList",
		MethodProperties: map[string]string{
			"GetFullList":  "1",
			"DateTimeFrom": "01.01.2021",
			"DateTimeTo":   "01.01.2050",
		},
		ApiKey: n.apiKey,
	}
	requestBodyBytes, _ := json.Marshal(requestBody)
	resp, apiError := http.Post(requestUrl, "application/json", bytes.NewReader(requestBodyBytes))
	if apiError != nil {
		return nil, apiError
	}
	var response GetInternetDocumentsResponse
	jsonError := json.NewDecoder(resp.Body).Decode(&response)
	if jsonError != nil {
		return nil, jsonError
	}
	if response.Success == true {
		return response.Data, nil
	} else {
		return nil, errors.New("failed to fetch internet documents")
	}
}

func (n NewPostServiceImpl) FilterInternetDocumentsAwaitingPickup(documents []InternetDocument) []InternetDocument {
	var awaitingPickup []InternetDocument
	for currentIndex := range documents {
		document := documents[currentIndex]
		documentDeliveryStatus := document.DeliveryStatus
		if documentDeliveryStatus == "7" || documentDeliveryStatus == "8" {
			awaitingPickup = append(awaitingPickup, document)
		}
	}
	return awaitingPickup
}

func (n NewPostServiceImpl) FilterDispatchedInternetDocuments(documents []InternetDocument) []InternetDocument {
	var awaitingPickup []InternetDocument
	for currentIndex := range documents {
		document := documents[currentIndex]
		documentDeliveryStatus := document.DeliveryStatus
		if documentDeliveryStatus == "1" {
			awaitingPickup = append(awaitingPickup, document)
		}
	}
	return awaitingPickup
}
