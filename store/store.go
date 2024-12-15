package store

type DataStore[K comparable, V any] interface {
	Set(key K, val V)
	Get(key K) (V, error)
	Delete(key K)
}
