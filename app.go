package main

import (
	"car_rent/routes"
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	r := routes.NewRouter()
	fmt.Printf("\nListenning on port: %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
