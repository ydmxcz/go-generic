package iterator

import (
	"github.com/ydmxcz/go-generic/collections/truple"
)

func Map[K comparable, V any](m map[K]V) Iter[truple.KV[K, V]] {
	s := make([]truple.KV[K, V], len(m))
	i := 0
	for k, v := range m {
		s[i].Key = k
		s[i].Val = v
		i++
	}
	return Slice(s)

}
