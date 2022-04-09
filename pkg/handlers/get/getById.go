package get

import (
	"net/http"

	"github.com/rahulsidpatil/mymdb/pkg/entities"
	"github.com/rahulsidpatil/mymdb/pkg/utils"
)

// @Summary Fetch movie by ID
// @Description Fetch movie by ID
// @Accept  json
// @Produce  json
// @Param id path int true "Movie ID"
// @Success 200 {object} entities.Movie
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /movies/{id} [get]
func GetById(w http.ResponseWriter, r *http.Request) {
	m := entities.Movie{
		Id:    "tt2884018",
		Title: "Avengers:Endgame",
		Year:  2019,
		Rated: 8.4,
		Genre: []string{"Action", "Adventure", "Drama"},
	}
	utils.RespondWithJSON(w, http.StatusOK, m)
}
