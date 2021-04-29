package router

import (
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/flashguru-git/node-monitor-server/controllers/api"
	_ "github.com/flashguru-git/node-monitor-server/docs"
	"github.com/flashguru-git/node-monitor-server/models"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if match, _ := regexp.MatchString("/swagger/*", r.URL.Path); match {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func NewRouter() *mux.Router {
	routes := mux.NewRouter()

	routes.Use(loggingMiddleware)
	routes.Use(commonMiddleware)

	routes.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/swagger/doc.json"), //The url pointing to API definition"
	))

	//append applications routes
	models.Routes = append(models.Routes, api.Routes)

	for _, route := range models.Routes {
		//create subroute
		routePrefix := routes.PathPrefix(route.Prefix).Subrouter()
		//loop through each sub route
		for _, r := range route.SubRoutes {
			var handler http.Handler
			handler = r.HandlerFunc
			//check to see if route should be protected with jwt
			if r.Protected {
				// authorization check
				// consider later
			}
			//attach sub route
			routePrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)
		}
	}
	return routes
}
