package main

import (
	"fmt"
	"net/http"
	"os"
)

type URLConfig struct {
	URL string
	Username string
	Password string
	Pattern string
}

const slackWebhookURL = "http://"

func processURL (usr string) error {
	

}

func main() {
	
	urls := []URLConfig {
		
		{
			URL:      "http://127.0.0.1:5500/Frontend/index.html",
			Username: "admin",
			Password: "admin123",
			Pattern:  `Cloudspace Metrics collected at (\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`,
		},
		{
			URL:      "http://127.0.0.1:5500/Frontend/report.html",
			Username: "admin",
			Password: "admin123",
			Pattern:  `Last Updated:\s+(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`,
		},
		{
			URL:      "http://127.0.0.1:5500/Frontend/stats.html",
			Username: "admin",
			Password: "admin123",
			Pattern:  `Report file last update on:\s+(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`,
		},
	}





}