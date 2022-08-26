package handlers

import (
	"github.com/bootcamp/website-status-checker/database"
	"github.com/bootcamp/website-status-checker/status_checker"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestViewWebsitesHandler(t *testing.T) {
	Initialize()
	req := httptest.NewRequest(http.MethodGet, "/view-websites-status", nil)
	w := httptest.NewRecorder()
	ViewWebsitesStatusHandler(w, req)
	res := w.Result()
	if res.StatusCode != 200 {
		log.Fatal("Test Failed")
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Error!")
		}
	}(res.Body)
	data, _ := ioutil.ReadAll(res.Body)
	expected := "{\"https://www.airbnb.com\":\"UP\",\"https://www.amazon.com\":\"UP\",\"https://www.google.com\":\"UP\"}"
	log.Printf("Expected:\t%v", expected)
	log.Printf("Output:\t\t%v", string(data))
}

func TestAddWebsitesHandler(t *testing.T) {
	reader := strings.NewReader("[{\"Url\":\"https://www.google.com\"}]")
	req := httptest.NewRequest(http.MethodPost, "/add-websites", reader)
	w := httptest.NewRecorder()
	AddWebsitesHandler(w, req)
	res := w.Result()
	if res.StatusCode != 200 {
		log.Fatal("Test Failed")
		return
	}
}

func Initialize() {
	MockWebsiteData()
	go status_checker.StatusCheckerGoRoutine()
	time.Sleep(1 * time.Second)
}

func MockWebsiteData() {
	database.WebsiteData = map[string]string{
		"https://www.airbnb.com": "DOWN",
		"https://www.amazon.com": "DOWN",
		"https://www.google.com": "DOWN",
	}
}
