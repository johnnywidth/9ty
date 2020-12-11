package repository

// DB db interface for set and get value for given key
type DB interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}
