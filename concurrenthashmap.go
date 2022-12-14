package concurrent

const defaultCapacity = 16
const defaultLoadFactor = 0.75
const defaultConcurrencyLevel = 1

type HashMap[K comparable, V any] struct {
	initialCapacity  uint
	loadFactor       float32
	concurrencyLevel uint

	/*
	 * The array of bins. Lazily initialized upon first insertion.
	 * Size is always a power of two. Accessed directly by iterators.
	 */
	table *node[K, V] // TODO ensure volatile semantics
	/*
	 * The next table to use; non-null only while resizing.
	 */
	nextTable *node[K, V] // TODO ensure volatile semantics
}

type node[K comparable, V any] struct {
	// TODO hash int
	// TODO key  K
	// TODO val  V
	// TODO next *node[K, V]
}

type Entry[K any, V any] struct {
	Key   K
	Value V
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
		table:            nil,
		nextTable:        nil,
	}
}

func (receiver *HashMap[K, V]) Size() uint {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) IsEmpty() bool {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) ContainsKey(key K) bool {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) ContainsValue(value V) bool {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) Put(key K, value V) V {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) Remove(key K) V {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) RemoveEntry(key K, value V) V {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) Replace(key K, value V) V {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) PutAll(hashMap HashMap[K, V]) {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) Clear() {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) KeySet() []K {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) ValueSet() []V {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) EntrySet() []Entry[K, V] {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) PutIfAbsent(key K, value V) V {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) ReplaceAll(function func(K, V) V) {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) ComputeIfAbsent(key K, function func(K) V) V {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) ComputeIfPresent(key K, remappingFunction func(K, V) V) V {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) Compute(key K, remappingFunction func(K, V) V) V {
	panic("TODO") // TODO
}

func (receiver *HashMap[K, V]) Merge(key K, value V, remappingFunction func(K, V) V) V {
	panic("TODO") // TODO
}
