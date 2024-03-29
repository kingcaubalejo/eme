package routers

import (
	"github.com/gorilla/mux"
)

func LoadRouter() *mux.Router {
	router := mux.NewRouter()

	router = usersRoute(router)
	router = rolesRoute(router)
	router = pastorsRoute(router)
	router = churchesRoute(router)
	
	router = auth(router)
	
	return router
}
