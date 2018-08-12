package pack
import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api",
		Index,
	},
	Route{
		"ShowAll",
		"GET",
		"/api/movies",
		TodoShowAll,
	},
	Route{
		"TodoShow",
		"GET",
		"/api/movies/{todoId}",
		TodoShowById,
	},
	Route{
		"TodoCreate",
		"POST",
		"/api/movie/add",
		TodoCreate,
	},
}