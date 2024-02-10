package routers

import (
	// "api/authorization"
	c "api/controllers/churches"

	"github.com/gorilla/mux"
	// "github.com/urfave/negroni"
)

func churchesRoute(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/church/add", c.Create).Methods("POST")
	router.HandleFunc("/v1/church/update/{churchId}", c.Update).Methods("PUT")
	router.HandleFunc("/v1/church/delete/{churchId}", c.Delete).Methods("DELETE")
	router.HandleFunc("/v1/churches", c.Get).Methods("GET")
	router.HandleFunc("/v1/church/{churchId}", c.GetInfo).Methods("GET")
	router.HandleFunc("/v1/church/{churchId}/pastor", c.WithPastor).Methods("GET")
	return router
}
