package firestore

import (
	"fmt"
	"log"
	"os"

	"github.com/alejandrowaiz98/Golang-Middlewares/entitys"
	"golang.org/x/net/context"
)

func (f *Firestore) Save(data []entitys.Billet) error {

	ctx := context.Background()

	for _, novel := range data {

		_, _, err := f.client.Collection(os.Getenv("Firestore-Collections-name")).Add(ctx, novel)

		if err != nil {
			log.Printf(("Firestore | ERROR) Err: %v"), err)
			return err
		}

		fmt.Println("¡Data saved into database!")

	}

	return nil

}
