package concurrent

type HashMap[K comparable, V any] struct {
}

func New[K comparable, V any]() HashMap[K, V] {
	return HashMap[K, V]{}
}

func NewWithInitialSize[K comparable, V any](size uint) HashMap[K, V] {
	return HashMap[K, V]{}
}

func NewWithLoadFactor[K comparable, V any](size uint, load_factor float32) HashMap[K, V] {
	return HashMap[K, V]{}
}
