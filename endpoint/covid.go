package endpoint

import (
	handler "covid-summary/business/covid"
	"covid-summary/business/covid/repository"
	"covid-summary/business/covid/service"
	"covid-summary/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func InitCovidEndpoint(router *gin.Engine, client *resty.Client) {
	covidRepo := repository.NewCovidRepository(client)
	covidSvc := service.NewCovidService(covidRepo)
	covidHandler := handler.NewCovidHandler(covidSvc)

	routeCovid := router.Group("/covid", middleware.Logger()) // case inject middleware
	routeCovid.GET("/summary", covidHandler.GetCovidSummary)
}
