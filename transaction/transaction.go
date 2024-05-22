package transaction

import (
	"suraj_projects/flipkart_machine_coding/utils"
	"sync"
)

type Transaction struct {
	TransactionId string
	OrderId       string
	Amount        int
	Commission    int
}

type TransMgr struct {
	TransMap map[string]*Transaction
}

var TransMgrInstance *TransMgr
var mu sync.Mutex

func GetTransMgr() *TransMgr {
	if TransMgrInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if TransMgrInstance == nil {
			TransMgrInstance = &TransMgr{
				TransMap: make(map[string]*Transaction),
			}
		}
	}
	return TransMgrInstance
}

func (t *TransMgr) CreateTransaction(txn Transaction) error {
	mu.Lock()
	defer mu.Unlock()
	id := utils.GenerateProductUniqueId()
	t.TransMap[id] = &txn
	return nil

}
