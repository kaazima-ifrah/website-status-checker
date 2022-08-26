package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bootcamp/website-status-checker/database"
	"github.com/bootcamp/website-status-checker/errors"
	"log"
	"net/http"
)

func AnonymousDataHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := fmt.Fprint(w, "\nInvalid! Please hit a valid request!\nHere are a few valid requests:\n/add-websites\n/view-websites-status\n\n")
	if err != nil {
		errors.HandleResponseWriterError("AnonymousDataHandler", err)
	}
}

func AddWebsitesHandler(w http.ResponseWriter, r *http.Request) {
	var websites []database.Website
	if err := json.NewDecoder(r.Body).Decode(&websites); err != nil {
		log.Println("Error while decoding the request body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Add websites to db & Initialize each website's status
	for _, website := range websites {
		database.WebsiteData[website.Url] = "DOWN"
	}

	_, err := json.Marshal(websites)
	if err != nil {
		log.Println("Could not marshal websites data:", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		_, err := fmt.Fprintf(w, "\nWebsites data is successfully stored!\n%+v\n\n", websites)
		if err != nil {
			errors.HandleResponseWriterError("AddWebsitesHandler", err)
		}
	}
}

func ViewWebsitesStatusHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("name")
	if queryParam != "" {
		ViewWebsite(w, queryParam)
	} else {
		ViewAllWebsites(w)
	}
}
