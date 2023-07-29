package service_test

import (
	"covid-summary/business/covid/mocks"
	"covid-summary/business/covid/service"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

var province = map[string]int{
	"Amnat Charoen":            17,
	"Ang Thong":                36,
	"Bangkok":                  27,
	"Bueng Kan":                23,
	"Buriram":                  18,
	"Chachoengsao":             24,
	"Chai Nat":                 25,
	"Chaiyaphum":               28,
	"Chanthaburi":              17,
	"Chiang Mai":               22,
	"Chiang Rai":               15,
	"Chonburi":                 29,
	"Chumphon":                 25,
	"Kalasin":                  27,
	"Kamphaeng Phet":           23,
	"Kanchanaburi":             23,
	"Khon Kaen":                27,
	"Krabi":                    27,
	"Lampang":                  24,
	"Lamphun":                  25,
	"Loei":                     17,
	"Lopburi":                  19,
	"Mae Hong Son":             22,
	"Maha Sarakham":            26,
	"Mukdahan":                 28,
	"N/A":                      27,
	"Nakhon Nayok":             19,
	"Nakhon Pathom":            31,
	"Nakhon Phanom":            24,
	"Nakhon Ratchasima":        28,
	"Nakhon Sawan":             24,
	"Nakhon Si Thammarat":      35,
	"Nan":                      20,
	"Narathiwat":               22,
	"Nong Bua Lamphu":          29,
	"Nong Khai":                27,
	"Nonthaburi":               29,
	"Pathum Thani":             30,
	"Pattani":                  27,
	"Phang Nga":                28,
	"Phatthalung":              29,
	"Phayao":                   25,
	"Phetchabun":               33,
	"Phetchaburi":              26,
	"Phichit":                  21,
	"Phitsanulok":              24,
	"Phra Nakhon Si Ayutthaya": 25,
	"Phrae":                    28,
	"Phuket":                   25,
	"Prachinburi":              19,
	"Prachuap Khiri Khan":      34,
	"Ranong":                   35,
	"Ratchaburi":               21,
	"Rayong":                   25,
	"Roi Et":                   25,
	"Sa Kaeo":                  26,
	"Sakon Nakhon":             42,
	"Samut Prakan":             31,
	"Samut Sakhon":             29,
	"Samut Songkhram":          22,
	"Saraburi":                 26,
	"Satun":                    37,
	"Sing Buri":                26,
	"Sisaket":                  27,
	"Songkhla":                 24,
	"Sukhothai":                23,
	"Suphan Buri":              28,
	"Surat Thani":              25,
	"Surin":                    24,
	"Tak":                      18,
	"Trang":                    20,
	"Trat":                     25,
	"Ubon Ratchathani":         23,
	"Udon Thani":               34,
	"Uthai Thani":              24,
	"Uttaradit":                24,
	"Yala":                     27,
	"Yasothon":                 26,
}

var ageGroup = map[string]int{
	"0-30":  602,
	"31-60": 607,
	"61+":   769,
	"N/A":   22,
}

// test happy case
func TestGetCovidSummary(t *testing.T) {
	var sumProvince, sumAgeGroup, sumNullProvince, sumNullAgeGroup int

	repo := &mocks.Repository{}
	repo.On("GetCovidSummary").Return(province, ageGroup, nil)

	svc := service.NewCovidService(repo)
	res, err := svc.GetCovidSummary()

	assert.NoError(t, err)

	for key, value := range res.Province {
		if key == "N/A" {
			sumNullProvince = value
		} else {
			sumProvince += value
		}
	}
	for key, value := range res.AgeGroup {
		if key == "N/A" {
			sumNullAgeGroup = value
		} else {
			sumAgeGroup += value
		}
	}

	assert.Equal(t, sumProvince, 1973)   // other provinces: 1973
	assert.Equal(t, sumAgeGroup, 1978)   // other age groups: 1978
	assert.Equal(t, sumNullProvince, 27) // N/A province: 27
	assert.Equal(t, sumNullAgeGroup, 22) // N/A age group: 22
}

// test error case
func TestGetCovidSummaryError(t *testing.T) {
	expectedErr := errors.New("error")

	repo := &mocks.Repository{}
	repo.On("GetCovidSummary").Return(nil, nil, expectedErr)

	svc := service.NewCovidService(repo)
	_, err := svc.GetCovidSummary()

	assert.Error(t, err)
	assert.EqualError(t, err, "error")
	assert.ErrorIs(t, err, expectedErr)
}

// test nil case
func TestGetCovidSummaryEmpty(t *testing.T) {
	repo := &mocks.Repository{}
	repo.On("GetCovidSummary").Return(map[string]int(nil), map[string]int(nil), nil)

	svc := service.NewCovidService(repo)
	res, err := svc.GetCovidSummary()

	assert.NoError(t, err)
	assert.Empty(t, res.Province)
	assert.Empty(t, res.AgeGroup)
}
