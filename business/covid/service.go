package covid

import (
	"context"
	"covid-summary/business/model"
)

type Service interface {
	GetCovidSummary(ctx context.Context) (model.ResponseCovidSummary, error)
}
