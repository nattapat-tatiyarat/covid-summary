package repository

import (
	"covid-summary/business/covid"
	"covid-summary/business/models"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type CovidRepository struct {
	client *resty.Client
}

func NewCovidRepository(client *resty.Client) covid.Repository {
	return &CovidRepository{
		client: client,
	}
}

func (r *CovidRepository) GetCovidSummary() (models.ResponseCovidCases, error) {
	result := models.ResponseCovidCases{}

	resp, err := r.client.R().SetResult(&result).Get("/devinterview/covid-cases.json")
	if err != nil {
		return models.ResponseCovidCases{}, err
	}
	if resp.IsError() {
		return models.ResponseCovidCases{}, fmt.Errorf("%s %d : %v", resp.Request.Method, resp.StatusCode(), resp)
	}

	return result, nil
}
