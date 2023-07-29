package models

type CovidCase struct {
	Age      *int    `json:"Age"`
	Province *string `json:"Province"`
}

type ResponseCovidCases struct {
	Data []CovidCase `json:"Data"`
}

type ResponseCovidSummary struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}
