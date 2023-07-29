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

func (r *CovidRepository) GetCovidSummary() (map[string]int, map[string]int, error) {
	result := models.ResponseCovidCases{}
	provinceMap := make(map[string]int)
	ageGroupMap := make(map[string]int)

	resp, err := r.client.R().SetResult(&result).Get("/devinterview/covid-cases.json")
	if err != nil {
		return nil, nil, err
	}
	if resp.IsError() {
		return nil, nil, fmt.Errorf("%d :: %v", resp.StatusCode(), resp)
	}

	for _, covidCase := range result.Data {
		// count Province
		if covidCase.Province != nil {
			provinceMap[*covidCase.Province]++
		} else {
			provinceMap["N/A"]++ // no province data
		}

		// count Age Group
		if covidCase.Age != nil {
			if 0 <= *covidCase.Age && *covidCase.Age <= 30 {
				ageGroupMap["0-30"]++ // between 0-30
			} else if 31 <= *covidCase.Age && *covidCase.Age <= 60 {
				ageGroupMap["31-60"]++ // between 31-60
			} else {
				ageGroupMap["61+"]++ // more than 60
			}
		} else {
			ageGroupMap["N/A"]++ // no age data
		}
	}

	return provinceMap, ageGroupMap, nil
}
