package main

import (
	"github.com/alejandrowaiz98/Golang-Middlewares/http"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	server, err := http.New()

	if err != nil {
		panic(err)
	}

	server.Run()

}
