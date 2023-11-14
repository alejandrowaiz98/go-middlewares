package service

import (
	"fmt"
	"log"

	"github.com/alejandrowaiz98/Golang-Middlewares/entitys"
)

var iva_Argentina float64
var iva_Chile float64

func (s *Service) Save(b []entitys.Billet) []error {

	billets, errors := applyIva(b)

	if len(errors) > 0 {
		log.Printf("Founded %v errors when trying to apply iva", len(errors))
		return errors
	}

	errors = s.db.Save(billets)

	if len(errors) > 0 {
		log.Printf("Founded %v errors when trying to save into database", len(errors))
		return errors
	}

	return nil

}

func applyIva(billets []entitys.Billet) ([]entitys.Billet, []error) {

	var errors []error
	var iva float64

	for _, b := range billets {

		switch b.Origin_Country {

		case "Argentina":

			for _, p := range b.Products {
				iva = p.Original_Price * (iva_Argentina * 100)
				p.Final_Price = p.Original_Price + iva
			}

		case "Chile":

			for _, p := range b.Products {
				iva = p.Original_Price * (iva_Chile * 100)
				p.Final_Price = p.Original_Price + iva
			}

		default:

			err := fmt.Errorf("[Service | Save | ApplyIva] error: country %v not available", b.Origin_Country)
			errors = append(errors, err)

		}

	}

	if len(errors) > 0 {
		return nil, errors
	}

	return billets, nil

}
