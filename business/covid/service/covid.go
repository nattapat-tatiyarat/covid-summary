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
	covidSummary, err := s.repo.GetCovidSummary()
	if err != nil {
		fmt.Println("Error GetCovidSummary : ", err.Error())
		return models.ResponseCovidSummary{}, err
	}

	province, ageGroup := make(map[string]int), make(map[string]int)

	for _, covidCase := range covidSummary.Data {
		// count Province
		if covidCase.Province != nil {
			province[*covidCase.Province]++
		} else {
			// no province data
			province["N/A"]++
		}

		// count Age Group
		if covidCase.Age != nil {
			if 0 <= *covidCase.Age && *covidCase.Age <= 30 {
				// between 0-30
				ageGroup["0-30"]++
			} else if 31 <= *covidCase.Age && *covidCase.Age <= 60 {
				// between 31-60
				ageGroup["31-60"]++
			} else {
				// more than 60
				ageGroup["61+"]++
			}
		} else {
			// no age data
			ageGroup["N/A"]++
		}
	}

	res := models.ResponseCovidSummary{
		Province: province,
		AgeGroup: ageGroup,
	}

	return res, nil
}
