package service

import (
	"log"

	"github.com/alejandrowaiz98/Golang-Middlewares/entitys"
	"github.com/alejandrowaiz98/Golang-Middlewares/firestore"
)

type Service struct {
	db firestore.FirestoreImplementation
}

type ServiceImplementation interface {
	Save(b []entitys.Billet) []error
}

func New() (ServiceImplementation, error) {

	db, err := firestore.New()

	if err != nil {
		log.Printf("Err creating service: %v", err)
		return nil, err
	}

	return &Service{db: db}, nil

}
