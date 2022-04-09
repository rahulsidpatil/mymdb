package entities

type Movie struct {
	Id    string   `json:"ImdbID,omitempty" example:"tt2884018"`
	Title string   `json:"Title,omitempty" example:"Avengers:Endgame"`
	Year  int      `json:"Year,omitempty" example:"2019"`
	Rated float32  `json:"Rated,omitempty" example:"8.4"`
	Genre []string `json:"Genre,omitempty" example:"Action, Adventure, Drama"`
}
