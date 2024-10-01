package mocks

type AerospikeClient interface {
	Put(binName string, value interface{}) error
	Get(binName string) (interface{}, error)
	Delete() error
}
