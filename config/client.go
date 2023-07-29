package config

import (
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

func InitClient() *resty.Client {
	baseUrl := os.Getenv("INTERVIEW_ENDPOINT")

	client := resty.New()
	client.SetBaseURL(baseUrl)
	client.SetTimeout(5 * time.Second)

	return client
}
