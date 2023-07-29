package covid

import "covid-summary/business/model"

type Repository interface {
	GetCovidSummary() (model.ResponseCovidCases, error)
}
