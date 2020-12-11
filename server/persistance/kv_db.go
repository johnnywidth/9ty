package persistance

import "sync"

type KvDB struct {
	mutex *sync.RWMutex
	data  map[string]interface{}
}

func NewKvDB() *KvDB {
	return &KvDB{
		mutex: &sync.RWMutex{},
		data:  make(map[string]interface{}),
	}
}

func (p *KvDB) Set(key string, value interface{}) error {
	p.mutex.Lock()
	p.data[key] = value
	p.mutex.Unlock()

	return nil
}

func (p *KvDB) Get(key string) (interface{}, error) {
	p.mutex.RLock()
	value := p.data[key]
	p.mutex.RUnlock()

	return value, nil
}
