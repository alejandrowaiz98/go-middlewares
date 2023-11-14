package service

import (
	"fmt"

	"github.com/alejandrowaiz98/Golang-Middlewares/entitys"
)

var iva_Argentina float64
var iva_Chile float64

func Save(b []entitys.Billet) {

}

func applyIva(billets []entitys.Billet) []error {

	var errors []error

	for _, b := range billets {

		switch b.Origin_Country {

		case "Argentina":

			for _, p := range b.Products {
				iva := p.Original_Price * (iva_Argentina * 100)
				p.Final_Price = p.Original_Price + iva
			}

		case "Chile":

			for _, p := range b.Products {
				iva := p.Original_Price * (iva_Chile * 100)
				p.Final_Price = p.Original_Price + iva
			}

		default:

			err := fmt.Errorf("[Service | Save | ApplyIva] error: country %v not available", b.Origin_Country)
			errors = append(errors, err)

		}

	}

	//TODO: Add firestore logic to add to database

	return nil

}
