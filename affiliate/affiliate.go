package affiliate

import (
	"fmt"
	"suraj_projects/flipkart_machine_coding/order"
	"suraj_projects/flipkart_machine_coding/utils"
	"sync"
)

type Affiliate struct {
	AffiliateId   string
	AffiliateName string
	Orders        []order.Order
}

type AffiliateMgr struct {
	AffiliateMap map[string]*Affiliate
	mu           *sync.Mutex
}

var mu sync.Mutex
var AffiliateMgrInstance *AffiliateMgr
var Affiliates = make(map[string]Affiliate)

func NewAffiliateMgr() *AffiliateMgr {
	if AffiliateMgrInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if AffiliateMgrInstance == nil {
			AffiliateMgrInstance = &AffiliateMgr{
				AffiliateMap: make(map[string]*Affiliate),
				mu:           &sync.Mutex{},
			}
		}

	}
	return AffiliateMgrInstance
}

func (am *AffiliateMgr) CreateAffiliate(affiliate Affiliate) error {
	am.mu.Lock()
	defer am.mu.Unlock()
	affiliateId := utils.GenerateAffiliateUniqueId()
	am.AffiliateMap[affiliateId] = &affiliate
	Affiliates[affiliateId] = affiliate
	return nil
}

func (am *AffiliateMgr) GetAffiliate(affiliateId string) (Affiliate, error) {
	am.mu.Lock()
	defer am.mu.Unlock()
	affiliate, ok := am.AffiliateMap[affiliateId]
	if !ok {
		return Affiliate{}, fmt.Errorf("Affiliate with id %s not found", affiliateId)
	}
	return *affiliate, nil
}

func (am *AffiliateMgr) GetAffiliateOrders(affiliateId string) ([]order.Order, error) {
	am.mu.Lock()
	defer am.mu.Unlock()
	affiliate, ok := am.AffiliateMap[affiliateId]

	if !ok {
		return nil, fmt.Errorf("Affiliate with id %s not found", affiliateId)
	}
	return affiliate.Orders, nil
}
