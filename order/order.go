package order

type Order struct {
	OrderId       string
	Price         int
	ProductId     string
	CategoryId    string
	AffiliateId   string
	CommisionType string
	Timestamp     string

	State string
}

var Orders = make(map[string]*Order)
