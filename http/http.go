package http

import (
	"log"

	"github.com/alejandrowaiz98/Golang-Middlewares/service"
)

type Http struct {
	s service.ServiceImplementation
}

type HttpImplementation interface {
}

func New() (HttpImplementation, error) {

	s, err := service.New()

	if err != nil {
		log.Printf("Err creating http client: %v", err)
		return nil, err
	}

	return Http{s: s}, nil
}
