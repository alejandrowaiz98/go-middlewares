package http

import (
	"github.com/alejandrowaiz98/Golang-Middlewares/firestore"
)

type Http struct {
	db firestore.FirestoreImplementation
}

type HttpImplementation interface {
}

func New(db firestore.FirestoreImplementation) HttpImplementation {

	return Http{db: db}
}
