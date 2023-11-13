package entitys

type Billet struct {
	TotalPrice int
	Products   []Product
	Country    string
}

type Product struct {
	Name, Description string
	Price             int
}
