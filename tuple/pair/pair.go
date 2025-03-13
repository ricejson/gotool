package pair

import "fmt"

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func (pair *Pair[K, V]) String() string {
	return fmt.Sprintf("<%#v, %#v>", pair.Key, pair.Value)
}

func (pair *Pair[K, V]) Split() (K, V) {
	return pair.Key, pair.Value
}

func NewPair[K any, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{
		Key:   key,
		Value: value,
	}
}
