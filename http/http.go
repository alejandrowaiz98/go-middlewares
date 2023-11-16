package http

import (
	"log"

	"github.com/alejandrowaiz98/Golang-Middlewares/service"
)

type Http struct {
	service service.ServiceImplementation
}

type HttpImplementation interface {
	Run()
}

func New() (HttpImplementation, error) {

	s, err := service.New()

	if err != nil {
		log.Printf("Err creating http client: %v", err)
		return nil, err
	}

	return &Http{service: s}, nil
}
