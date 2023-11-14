package entitys

type Billet struct {
	TotalPrice     float64
	Products       []Product
	Origin_Country string
}

type Product struct {
	Name, Description           string
	Original_Price, Final_Price float64
}
