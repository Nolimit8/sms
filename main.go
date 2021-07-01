package main

import (
	"os"
)

func main() {
	newPostApiKey := os.Getenv("NP_API_KEY")
	newPostService := NewPostServiceImpl{
		apiKey: newPostApiKey,
	}
	resp, err := newPostService.GetAllInternetDocuments()
	if err != nil {
		print(resp)
	} else {
		print(err)
	}
}
