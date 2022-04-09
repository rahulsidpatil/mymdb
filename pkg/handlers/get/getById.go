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
		Id:          1,
		Title:       "Avengers:Endgame",
		ReleaseYear: 2019,
		Rating:      8.4,
		Genere:      []string{"Action", "Adventure", "Drama"},
		MetaData: `After the devastating events of Avengers: Infinity War (2018), the universe is in ruins. 
		With the help of remaining allies, the Avengers assemble once more in order to reverse Thanos' 
		actions and restore balance to the universe.`,
	}
	utils.RespondWithJSON(w, http.StatusOK, m)
}
