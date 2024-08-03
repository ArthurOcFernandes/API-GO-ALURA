package routes

import (
	"api-go-rest/controller"
	"api-go-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/api/personalidades", controller.Personalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controller.PersonalidadePorID).Methods("Get")
	r.HandleFunc("/api/personalidades", controller.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controller.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controller.EditaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
