package firestore

import (
	"fmt"
	"log"
	"os"

	"github.com/alejandrowaiz98/Golang-Middlewares/entitys"
	"golang.org/x/net/context"
)

var collection string = os.Getenv("Firestore-Collections-name")

func (f *Firestore) Save(data []entitys.Billet) []error {

	ctx := context.Background()
	var errors []error

	for _, billet := range data {

		_, _, err := f.client.Collection(collection).Add(ctx, billet)

		if err != nil {
			log.Printf(("Firestore | ERROR) Err: %v"), err)
			errors = append(errors, err)
		}

		fmt.Println("¡Data saved into database!")

	}

	if len(errors) == 0 {

		return nil

	}

	return errors

}
