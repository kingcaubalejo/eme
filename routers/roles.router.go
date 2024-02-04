package routers

import (
	// "api/authorization"
	c "api/controllers/roles"

	"github.com/gorilla/mux"
	// "github.com/urfave/negroni"
)

func rolesRoute(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/roles/add", c.Create).Methods("POST")
	router.HandleFunc("/v1/roles/update/{roleId}", c.Update).Methods("PUT")
	router.HandleFunc("/v1/roles/delete/{roleId}", c.Delete).Methods("DELETE")
	router.HandleFunc("/v1/roles", c.Get).Methods("GET")
	router.HandleFunc("/v1/roles/{roleId}", c.GetInfo).Methods("GET")
	return router
}
