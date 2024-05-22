package order

type CommissionStrategy interface {
	CalculateCommission(order Order) int
}

type Commission struct {
	CommissionId string
	OrderId      string
	Commission   int
}

type CommissionServiceRuleInterface interface {
	// CreateCommission creates a new commission
	CreateCommissionRule(commission Commission) error
}

type FlatCommission struct {
}

func (fc *FlatCommission) CalculateCommission(order Order) int {
	return 10
}

type PercentageCommission struct {
}

func (pc *PercentageCommission) CalculateCommission(order Order) int {
	return 5
}
