package concurrent

const defaultCapacity = 16
const defaultLoadFactor = 0.75
const defaultConcurrencyLevel = 1

const hashBits int = 0x7fffffff // usable bits of normal node hash

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

	/*
	 * Base counter value, used mainly when there is no contention,
	 * but also as a fallback during table initialization
	 * races. Updated via CAS.
	 */
	baseCount int64

	/*
	 * Table initialization and resizing control.  When negative, the
	 * table is being initialized or resized: -1 for initialization,
	 * else -(1 + the number of active resizing threads).  Otherwise,
	 * when table is null, holds the initial table size to use upon
	 * creation, or 0 for default. After initialization, holds the
	 * next element count value upon which to resize the table.
	 */
	sizeCtl int

	/*
	 * The next table index (plus one) to split while resizing.
	 */
	transferIndex int

	/*
	 * Spinlock (locked via CAS) used when resizing and/or creating CounterCells.
	 */
	cellsBusy int
}

type node[K comparable, V any] struct {
	hash int
	key  K
	val  V
	next *node[K, V]
}

func (r *node[K, V]) find(hash int, key *K) *node[K, V] {
	if key == nil {
		return nil
	}
	current := r
	if current.hash == hash && current.key == *key {
		return current
	}

	for current.next != nil {
		current = current.next
		if current.hash == hash && current.key == *key {
			return current
		}
	}

	return nil
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
