package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rahulsidpatil/mymdb/api/docs"
	"github.com/rahulsidpatil/mymdb/pkg/handlers/get"
	"github.com/rahulsidpatil/mymdb/pkg/utils"
	httpSwagger "github.com/swaggo/http-swagger"
)

var svcPort string

type MyMdb struct {
	router *mux.Router
}

func (m *MyMdb) Initialize() {
	svcVersion := os.Getenv("SVC_VERSION")
	svcPort = os.Getenv("SVC_PORT")
	if svcPort == "" {
		svcPort = "8080"
	}
	svcPathPrefix := svcVersion + "/" + os.Getenv("SVC_PATH_PREFIX")
	swaggerAddr := "http://localhost:" + svcPort + "/swagger/doc.json"

	m.router = mux.NewRouter()
	setSwaggerInfo(swaggerAddr, svcPort, svcVersion)
	m.initializeRoutes(swaggerAddr, svcPathPrefix)
}

func (m *MyMdb) Run() {
	svcAddr := ":" + svcPort
	log.Printf("starting message server at:%s", svcAddr)
	log.Println(http.ListenAndServe(svcAddr, m.router))
}

func (m *MyMdb) initializeRoutes(swaggerAddr, svcPathPrefix string) {
	m.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(swaggerAddr), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	//welcome to mymdb
	m.router.HandleFunc(os.Getenv("SVC_VERSION")+"/hello", hello).Methods("GET")
	m.router.HandleFunc(svcPathPrefix+"/{id:[0-9]+}", get.GetById).Methods("GET")
	m.router.HandleFunc(svcPathPrefix+"/{title:[a-z0-9]}", get.GetByTitle).Methods("GET")
}

func setSwaggerInfo(swaggerAddr, port, version string) {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample mymdb server."
	docs.SwaggerInfo.Version = "1.0"
	//TODO: remove hard-coding of Host address
	docs.SwaggerInfo.Host = "localhost:" + port
	docs.SwaggerInfo.BasePath = version
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

// @Summary Say hello to user
// @Description Say hello to user
// @Success 200 {object} string
// @Failure 404
// @Router /hello [get]
func hello(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, "Welcome to My Movies Database aka mymdb..!!!")
}
