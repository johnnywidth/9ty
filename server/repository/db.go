package repository

type DB interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}
