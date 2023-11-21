package firestore

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/alejandrowaiz98/Golang-Middlewares/entitys"
	"google.golang.org/api/option"
)

type Firestore struct {
	client     *firestore.Client
	collection string
}

type FirestoreImplementation interface {
	Save(data []entitys.Billet) []error
}

var col string

func New() (FirestoreImplementation, error) {

	ctx := context.Background()

	col = os.Getenv("firestore_collections_name")

	client, err := firestore.NewClient(ctx, os.Getenv("firestore_project_id"), option.WithCredentialsFile(os.Getenv("service_account_filepath")))

	if err != nil {
		err = fmt.Errorf("(Firestore | NewClient) Err: %v ", err)
		return nil, err
	}

	return &Firestore{client: client, collection: col}, nil

}
