package database

type Website struct {
	Url string `json:"url"`
}

var WebsiteData = make(map[string]string)
