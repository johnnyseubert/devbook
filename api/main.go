package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/johnnyseubert/devbook/src/router"
)

func main() {
	r := router.Generate()

	fmt.Println("API is running on http://localhost:8080 ❤")
	log.Fatal(http.ListenAndServe(":8080", r))
}
