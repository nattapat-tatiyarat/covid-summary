package service_test

import (
	"covid-summary/business/covid/mocks"
	"covid-summary/business/covid/service"
	"covid-summary/business/models"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

// happy case
func TestGetCovidSummary(t *testing.T) {
	// mock province
	bangkok := "Bangkok"
	nonthaburi := "Nonthaburi"
	pathumThani := "Pathum Thani"

	// mock age
	childToAdult := 15
	adultToSenior := 45
	senior := 70

	repo := &mocks.Repository{}
	repo.On("GetCovidSummary").Return(models.ResponseCovidCases{
		Data: []models.CovidCase{
			{Province: &bangkok, Age: &childToAdult},
			{Province: &bangkok, Age: &childToAdult},
			{Province: &nonthaburi, Age: &adultToSenior},
			{Province: &nonthaburi, Age: &adultToSenior},
			{Province: &pathumThani, Age: &senior},
			{Province: &pathumThani, Age: &senior},
			{Province: nil, Age: nil},
			{Province: nil, Age: nil},
		},
	}, nil)

	svc := service.NewCovidService(repo)
	res, err := svc.GetCovidSummary()

	assert.NoError(t, err)

	assert.Equal(t, res.Province["Bangkok"], 2)
	assert.Equal(t, res.Province["Nonthaburi"], 2)
	assert.Equal(t, res.Province["Pathum Thani"], 2)
	assert.Equal(t, res.Province["N/A"], 2)

	assert.Equal(t, res.AgeGroup["0-30"], 2)
	assert.Equal(t, res.AgeGroup["31-60"], 2)
	assert.Equal(t, res.AgeGroup["61+"], 2)
	assert.Equal(t, res.AgeGroup["N/A"], 2)
}

// error case
func TestGetCovidSummaryError(t *testing.T) {
	expectedErr := errors.New("error")

	repo := &mocks.Repository{}
	repo.On("GetCovidSummary").Return(models.ResponseCovidCases{}, expectedErr)

	svc := service.NewCovidService(repo)
	_, err := svc.GetCovidSummary()

	assert.Error(t, err)
	assert.EqualError(t, err, "error")
	assert.ErrorIs(t, err, expectedErr)
}

// nil case
func TestGetCovidSummaryEmpty(t *testing.T) {
	repo := &mocks.Repository{}
	repo.On("GetCovidSummary").Return(models.ResponseCovidCases{}, nil)

	svc := service.NewCovidService(repo)
	res, err := svc.GetCovidSummary()

	assert.NoError(t, err)
	assert.Empty(t, res.Province)
	assert.Empty(t, res.AgeGroup)
}
