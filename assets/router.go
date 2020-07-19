package assets

import (
	"net/http"
	"context"
	"github.com/gorilla/mux"
	//logger "github.com/sirupsen/logrus"
)

//Get a router for handling your HTTP requests
func NewRouter(ctx context.Context, service Service) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		//handler = logger.LogHttpRequest(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
