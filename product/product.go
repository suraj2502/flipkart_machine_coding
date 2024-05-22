package product

import "sync"

type Product struct {
	ProductId      string
	CategoryId     string
	Price          int
	ProductName    string
	CommissionType string
}

type ProductMgr struct {
	ProductMap map[string]*Product
	mu         *sync.Mutex
}

var mu sync.Mutex

var ProductMgrInstance *ProductMgr

func NewProductMgr() *ProductMgr {
	if ProductMgrInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if ProductMgrInstance == nil {
			ProductMgrInstance = &ProductMgr{
				ProductMap: make(map[string]*Product),
				mu:         &sync.Mutex{},
			}
		}

	}
	return ProductMgrInstance
}

func (pm *ProductMgr) CreateProduct(product Product) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.ProductMap[product.ProductId] = &product
	return nil
}
