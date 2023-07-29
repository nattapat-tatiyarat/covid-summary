package covid

import "covid-summary/business/model"

type Service interface {
	GetCovidSummary() (model.ResponseCovidSummary, error)
}
