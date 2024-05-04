package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	// read the file .env
	godotenv.Load(".env") // read the file .env

	// use the variable PORT from the .evn file
	portString := os.Getenv("PORT")

	// if we get empty string we can not read the PORT variable form .env file.
	if portString == "" {
		log.Fatal("PORT is not found in the enviroment")
	}

	// creating a new mux object
	router := chi.NewRouter()

	//setting the CORS requests
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//Endpoints
	v1Router := chi.NewRouter()
	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	//  HTTP server with a specified router yo handle incoming requests and specifies the network address
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Printf("Server runs on port %v", portString)

	// error handling on the server , starts the server and if we have some error we get the error
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
