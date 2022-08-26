package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bootcamp/website-status-checker/database"
	"github.com/bootcamp/website-status-checker/errors"
	"log"
	"net/http"
)

func ViewAllWebsites(w http.ResponseWriter) {
	if len(database.WebsiteData) == 0 {
		_, err := fmt.Fprintf(w, "\nNo websites are available in the database :(\nPlease add websites to fetch the websites status!\n\n")
		if err != nil {
			errors.HandleResponseWriterError("ViewAllWebsites", err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(database.WebsiteData); err != nil {
			log.Println("Error while encoding the website data:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		_, err := fmt.Fprint(w)
		if err != nil {
			errors.HandleResponseWriterError("ViewAllWebsites", err)
		}
	}
}

func ViewWebsite(w http.ResponseWriter, queryParam string) {
	status := database.WebsiteData[queryParam]
	if status == "" {
		_, err := fmt.Fprintf(w, "\n'%v' website not found in the database!\nPlease add the website to find its status\n\n", queryParam)
		if err != nil {
			errors.HandleResponseWriterError("ViewWebsite", err)
		}
	}
	_, err := fmt.Fprintf(w, "\nThe status of the website '%v' is %v\n\n", queryParam, status)
	if err != nil {
		errors.HandleResponseWriterError("ViewWebsite", err)
	}
}
