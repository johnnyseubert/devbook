package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DatabaseConectionString string = ""
	Port                           = 0
)

func Load() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	DatabaseConectionString = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"

	show()
}

func show() {
	fmt.Println("===============================")
	fmt.Println("Variaveis de ambiente carregadas com sucesso")
	fmt.Println("[API_PORT]", Port)
	fmt.Println("[DB_CONNECTION_STRING]", DatabaseConectionString)
	fmt.Println("===============================")
}
