package concurrent

const defaultCapacity = 16
const defaultLoadFactor = 0.75
const defaultConcurrencyLevel = 1

type HashMap[K comparable, V any] struct {
	initialCapacity  uint
	loadFactor       float32
	concurrencyLevel uint
}

func New[K comparable, V any]() HashMap[K, V] {
	return NewWithInitialSize[K, V](defaultCapacity)
}

func NewWithInitialSize[K comparable, V any](capacity uint) HashMap[K, V] {
	return NewWithLoadFactor[K, V](capacity, defaultLoadFactor)
}

func NewWithLoadFactor[K comparable, V any](capacity uint, loadFactor float32) HashMap[K, V] {
	return NewWithConcurrencyLevel[K, V](capacity, loadFactor, defaultConcurrencyLevel)
}

func NewWithConcurrencyLevel[K comparable, V any](capacity uint, loadFactor float32, concurrencyLevel uint) HashMap[K, V] {
	return HashMap[K, V]{
		concurrencyLevel: concurrencyLevel,
		initialCapacity:  capacity,
		loadFactor:       loadFactor,
	}
}
