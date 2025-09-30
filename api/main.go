package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/johnnyseubert/devbook/src/config"
	"github.com/johnnyseubert/devbook/src/router"
)

func main() {
	config.Load()
	r := router.Generate()

	fmt.Println("API is running ❤")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
