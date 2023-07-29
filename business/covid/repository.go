package covid

type Repository interface {
	GetCovidSummary() (map[string]int, map[string]int, error)
}
