package main

import (
	"com.maxkucher/np-customer-reminder/job"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", job.RunJob)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
