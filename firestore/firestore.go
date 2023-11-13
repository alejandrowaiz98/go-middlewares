package firestore

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type Firestore struct {
	client *firestore.Client
}

type FirestoreImplementation interface {
}

func New() (FirestoreImplementation, error) {

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("Project-id"), option.WithCredentialsFile(os.Getenv("sa-filepath")))

	if err != nil {
		err = fmt.Errorf("(Firestore | NewClient) Err: %v ", err)
		return nil, err
	}

	return &Firestore{client: client}, nil

}