package routers

import (
	// "api/authorization"
	c "api/controllers/pastors"

	"github.com/gorilla/mux"
	// "github.com/urfave/negroni"
)

func pastorsRoute(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/pastors/add", c.Create).Methods("POST")
	router.HandleFunc("/v1/pastors/update/{pastorId}", c.Update).Methods("PUT")
	router.HandleFunc("/v1/pastors/delete/{pastorId}", c.Delete).Methods("DELETE")
	router.HandleFunc("/v1/pastors", c.Get).Methods("GET")
	router.HandleFunc("/v1/pastors/{pastorId}", c.GetInfo).Methods("GET")
	return router
}
