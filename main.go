package main

import (
	"covid-summary/config"
	"covid-summary/endpoint"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	// An Engine instance with the Logger and Recovery middleware already attached
	r := gin.Default()

	// init resty client
	client := config.InitClient()

	// handle no route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "NOT_FOUND", "message": "route not found"})
	})

	// set trusted proxies
	trustedProxies := []string{os.Getenv("INTERVIEW_IP_ADDRESS")}
	if err := r.SetTrustedProxies(trustedProxies); err != nil {
		log.Fatalf("set trusted proxies failed : %s", err.Error())
	}

	// init endpoint
	endpoint.InitCovidEndpoint(r, client)

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("start service failed : %s", err.Error())
	}
}
