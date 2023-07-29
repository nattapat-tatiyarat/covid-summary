package service

import (
	"covid-summary/business/covid"
	"covid-summary/business/models"
	"fmt"
)

type CovidService struct {
	repo covid.Repository
}

func NewCovidService(covidRepository covid.Repository) covid.Service {
	return &CovidService{
		repo: covidRepository,
	}
}

func (s *CovidService) GetCovidSummary() (models.ResponseCovidSummary, error) {
	province, ageGroup, err := s.repo.GetCovidSummary()
	if err != nil {
		fmt.Println("Error GetCovidSummary : ", err)
		return models.ResponseCovidSummary{}, err
	}

	res := models.ResponseCovidSummary{
		Province: province,
		AgeGroup: ageGroup,
	}

	return res, nil
}
