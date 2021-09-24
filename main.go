package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chattes/gta-schools-info/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("ENV FILE MISSING!")
	}
	router.SetupRoutes()
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	port = fmt.Sprintf(":%s", port)
	fmt.Printf("Listening on %s", port)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}

}
