package main

import (
	"github.com/flashguru-git/node-monitor-server/dao"
	"github.com/flashguru-git/node-monitor-server/router"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	dao.Connect()
	dao.PopulateIndex()
}

// Define HTTP request routes
func main() {
	godotenv.Load()
	routes := router.NewRouter()

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Println("Server Is Running At ", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(headersOk, originsOk, methodsOk)(routes)))
}
