package firestore

import (
	"fmt"
	"log"

	"github.com/alejandrowaiz98/Golang-Middlewares/entitys"
	"golang.org/x/net/context"
)

func (f *Firestore) Save(data []entitys.Billet) []error {

	ctx := context.Background()
	var errors []error

	log.Printf("data from frontend :%v", data)

	for _, billet := range data {

		_, _, err := f.client.Collection(f.collection).Add(ctx, &billet)

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
