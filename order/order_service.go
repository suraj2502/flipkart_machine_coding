package order

import (
	"fmt"
	"suraj_projects/flipkart_machine_coding/utils"
	"sync"
)

type OrderServiceInterface interface {
	// CreateOrder creates a new order
	CreateOrder(order Order) error
	// UpdateOrderState updates the state of an order
	UpdateOrderState(orderId string, state string) error
	// GetOrder returns the order details
	GetOrder(orderId string) (Order, error)
	// GetOrdersByStateAndAffiliateId returns all orders by state and affiliateId
	GetOrdersByStateAndAffiliateId(state string, affiliateId string) ([]Order, error)
	// // GetTransactionsByAffiliateId returns all transactions by affiliateId
	// GetTransactionsByAffiliateId(affiliateId string) ([]Transaction, error)
}

type OrderService struct {
	OrderMap map[string]*Order

	mu *sync.Mutex
}

func NewOrderService() *OrderService {
	return &OrderService{
		OrderMap: make(map[string]*Order),

		mu: &sync.Mutex{},
	}
}

func (os *OrderService) CreateOrder(order Order) error {
	os.mu.Lock()
	defer os.mu.Unlock()
	orderId := utils.GenerateOfferUniqueId()
	order.State = "Created"
	Orders[orderId] = &order

	if order.OrderId != "" {
		orderId = order.OrderId
	}
	os.OrderMap[orderId] = &order
	return nil
}

func (os *OrderService) UpdateOrderState(orderId string, state string) error {
	os.mu.Lock()
	defer os.mu.Unlock()
	order, ok := os.OrderMap[orderId]
	if !ok {
		return fmt.Errorf("Order with id %s not found", orderId)
	}
	order.State = state
	return nil
}

func (os *OrderService) GetOrder(orderId string) (Order, error) {
	os.mu.Lock()
	defer os.mu.Unlock()
	order, ok := os.OrderMap[orderId]
	if !ok {
		return Order{}, fmt.Errorf("Order with id %s not found", orderId)
	}
	return *order, nil
}

func (os *OrderService) GetOrdersByStateAndAffiliateId(state string, affiliateId string) ([]Order, error) {
	os.mu.Lock()
	defer os.mu.Unlock()
	var orders []Order
	for _, order := range os.OrderMap {
		if order.State == state && order.AffiliateId == affiliateId {
			orders = append(orders, *order)
		}
	}
	return orders, nil
}

func (os *OrderService) GetTotalCommisionForAffiliate(affiliateId string) (int, error) {
	os.mu.Lock()
	defer os.mu.Unlock()
	var totalCommision int
	for _, order := range os.OrderMap {
		if order.AffiliateId == affiliateId && order.State == "Delivered" && utils.IsCancelWindowCrossed(order.Timestamp, 2) {
			if order.CommisionType == "Percentage" {
				totalCommision += order.Price * 10 / 100
			} else {
				totalCommision += 5
			}
		}
	}
	return totalCommision, nil

}
