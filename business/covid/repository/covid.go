package repository

import (
	"covid-summary/business/covid"
	"covid-summary/business/model"
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

func (r *CovidRepository) GetCovidSummary() (model.ResponseCovidCases, error) {
	result := model.ResponseCovidCases{}

	resp, err := r.client.R().SetResult(&result).Get("/devinterview/covid-cases.json")
	if err != nil {
		return model.ResponseCovidCases{}, err
	}
	if resp.IsError() {
		return model.ResponseCovidCases{}, fmt.Errorf("%s %d : %v", resp.Request.Method, resp.StatusCode(), resp)
	}

	return result, nil
}
