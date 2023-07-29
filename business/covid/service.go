package covid

import "covid-summary/business/models"

type Service interface {
	GetCovidSummary() (models.ResponseCovidSummary, error)
}
