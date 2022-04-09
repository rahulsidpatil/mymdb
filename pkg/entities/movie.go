package entities

type Movie struct {
	Id          int         `json:"id,omitempty" example:"1"`
	Title       string      `json:"title,omitempty" example:"Avengers:Endgame"`
	ReleaseYear int         `json:"release_year,omitempty" example:"2019"`
	Rating      float32     `json:"rating,omitempty" example:"8.4"`
	Genere      []string    `json:"genere,omitempty" example:"[\"Action\",\"Adventure\",\"Drama\"]"`
	MetaData    interface{} `json:"metadata,omitempty"`
}
