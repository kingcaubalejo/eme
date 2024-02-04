package routers

import (
	// "api/authorization"
	c "api/controllers/users"

	"github.com/gorilla/mux"
	// "github.com/urfave/negroni"
)

func usersRoute(router *mux.Router) *mux.Router {

	router.HandleFunc("/v1/users/add", c.Create).Methods("POST")
	router.HandleFunc("/v1/users/update", c.Update).Methods("PUT")
	// router.Handle("/v1/users/delete",
	// 	negroni.New(
	// 		negroni.HandlerFunc(authorization.IsAuthorized),
	// 		negroni.HandlerFunc(c.Delete),
	// 	)).
	// 	Methods("DELETE")
	router.HandleFunc("/v1/users/get", c.Get).Methods("GET")
	router.HandleFunc("/v1/users/get-info/{id}", c.GetInfo).Methods("GET")

	return router
}
