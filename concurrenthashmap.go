package concurrent

const defaultCapacity = 16
const defaultLoadFactor = 0.75
const defaultConcurrencyLevel = 1

// treeify

const treeifyThreshold = 8
const untreeifyThreshold = 6
const minTreeifyCapacity = 64

type HashMap[K comparable, V any] struct {
}

func New[K comparable, V any]() HashMap[K, V] {
	return NewWithInitialSize[K, V](defaultCapacity)
}

func NewWithInitialSize[K comparable, V any](capacity uint) HashMap[K, V] {
	return NewWithLoadFactor[K, V](defaultCapacity, defaultLoadFactor)
}

func NewWithLoadFactor[K comparable, V any](capacity uint, load_factor float32) HashMap[K, V] {
	return NewWithConcurrencyLevel[K, V](defaultCapacity, defaultLoadFactor, defaultConcurrencyLevel)
}

func NewWithConcurrencyLevel[K comparable, V any](capacity uint, load_factor float32, concurrency_level uint) HashMap[K, V] {
	return HashMap[K, V]{}
}
