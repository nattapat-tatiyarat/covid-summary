package service

import (
	"context"
	"covid-summary/business/covid"
	"covid-summary/business/model"
	"fmt"

	"gorm.io/gorm/logger"
)

type CovidService struct {
	repo covid.Repository
}

func NewCovidService(covidRepository covid.Repository) covid.Service {
	return &CovidService{
		repo: covidRepository,
	}
}

func (s *CovidService) GetCovidSummary(ctx context.Context) (model.ResponseCovidSummary, error) {
	covidSummary, err := s.repo.GetCovidSummary()
	if err != nil {
		logger.Default.Error(ctx, fmt.Sprintf("GetCovidSummary failed : %s", err.Error()))
		return model.ResponseCovidSummary{}, err
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

	res := model.ResponseCovidSummary{
		Province: province,
		AgeGroup: ageGroup,
	}

	return res, nil
}
