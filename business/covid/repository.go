package covid

import "covid-summary/business/models"

type Repository interface {
	GetCovidSummary() (models.ResponseCovidCases, error)
}
