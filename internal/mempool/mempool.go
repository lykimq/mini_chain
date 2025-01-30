package mempool

import (
	"sync"
)

type Mempool struct {
	transactions []string
	mutex sync.RWMutex
}

func NewMempool() *Mempool {
	return &Mempool{
		transactions: make([]string, 0)
	}
}

func (mp *Mempool) AddTrasaction (data string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.transactions = append(m.transactions, data)
}

func (mp *Mempool) GetTransactions() []string {
	mp.mutex.RLock()
	defer mp.mutex.RUnlock()

	result := make([]string, len(mp.transactions))

	copy(result, mp.transactions)
	return result
}

func (mp *Mempool) Clear() {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()

	mp.transactions = make([]string, 0)
}