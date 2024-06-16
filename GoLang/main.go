package main

import (
	"fmt"
	"log"
	"mongo/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDb API")
	r := router.Router()
	fmt.Println("server getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listning at port 4000")
}
