package main

import (
	"github.com/bootcamp/website-status-checker/handlers"
	"github.com/bootcamp/website-status-checker/status_checker"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server at port 3000...")
	http.HandleFunc("/", handlers.AnonymousDataHandler)
	http.HandleFunc("/add-websites", handlers.AddWebsitesHandler)
	http.HandleFunc("/view-websites-status", handlers.ViewWebsitesStatusHandler)
	go status_checker.StatusCheckerGoRoutine()
	if err := http.ListenAndServe("127.0.0.1:3000", nil); err != nil {
		log.Println("Error listening to the server!")
		return
	}
}
