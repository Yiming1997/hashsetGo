package hashgo

type HashItem[T comparable, V any] struct {
	key T
	val V
}

func (hi *HashItem[T, V]) GetKey() T {
	return hi.key
}

func (hi *HashItem[T, V]) GetVal() V {
	return hi.val
}

func (hi *HashItem[T, V]) SetKey(key T) {
	hi.key = key
}
