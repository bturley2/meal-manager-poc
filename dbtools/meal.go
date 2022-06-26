package dbtools

const (
	chicken Protein = iota
	beef
	turkey
	pork
	fish
	veggie
	other
)

var (
	// treat as const
	acceptedProteins = []string{"chicken", "beef", "turkey", "pork", "fish", "veggie", "other"}
)

type Protein int

type Meal struct {
	Protein Protein `json:"Protein"`
	Title   string  `json:"title"`
	Url     string  `json:"url"`
	Rating  int     `json:"rating"`
	Notes   string  `json:"notes"`
}

func IsValidProtein(p string) bool {
	for _, v := range acceptedProteins {
		if p == v {
			return true
		}
	}
	return false
}

func StringToProtein(p string) Protein {
	for i, v := range acceptedProteins {
		if p == v {
			return Protein(i)
		}
	}
	return -1
}
