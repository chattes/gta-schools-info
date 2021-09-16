package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chattes/gta-schools-info/router"
)

func main() {
	router.SetupRoutes()
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	port = fmt.Sprintf(":%s", port)
	fmt.Printf("Listening on %s", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}

}
