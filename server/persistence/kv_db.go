package persistence

import "sync"

// KvDB Key value datastore with mutex and map of interfaces by string key
type KvDB struct {
	mutex *sync.RWMutex
	data  map[string]interface{}
}

// NewKvDB new instance of KV datastore
func NewKvDB() *KvDB {
	return &KvDB{
		mutex: &sync.RWMutex{},
		data:  make(map[string]interface{}),
	}
}

// Set set value for given key
func (p *KvDB) Set(key string, value interface{}) error {
	p.mutex.Lock()
	p.data[key] = value
	p.mutex.Unlock()

	return nil
}

// Get retrieve value for given key
func (p *KvDB) Get(key string) (interface{}, error) {
	p.mutex.RLock()
	value := p.data[key]
	p.mutex.RUnlock()

	return value, nil
}
