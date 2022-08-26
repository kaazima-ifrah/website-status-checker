package status_checker

import (
	"github.com/bootcamp/website-status-checker/database"
	"log"
	"net/http"
	"sync"
	"time"
)

func StatusCheckerGoRoutine() {
	for {
		var wg = sync.WaitGroup{}
		wg.Add(len(database.WebsiteData))
		log.Println("Checking status of websites...")
		for website := range database.WebsiteData {
			go func(website string) {
				log.Println("Checking status of ", website)
				status, err := CheckStatus(website)
				if err != nil {
					log.Fatal("Error while checking the website status: ", err)
					return
				}
				database.WebsiteData[website] = status
				log.Printf("Status of %v is %v\n", website, status)
				wg.Done()
			}(website)
		}
		wg.Wait()
		log.Println("Finished checking status of all websites, will recheck after a minute...")
		time.Sleep(1 * time.Second)
	}
}

func CheckStatus(url string) (string, error) {
	response, err := http.Get(url)
	if response != nil {
		status := RetrieveStatusFromStatusCode(response.StatusCode)
		return status, err
	}
	return "", err
}

func RetrieveStatusFromStatusCode(statusCode int) string {
	if statusCode == 200 {
		return "UP"
	}
	return "DOWN"
}
