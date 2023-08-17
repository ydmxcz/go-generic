package stream

import (
	"github.com/ydmxcz/go-generic/collections/truple"
	"github.com/ydmxcz/go-generic/iterator"
)

func Zip[T, R any](stm Stream[T], stmR Stream[R]) Stream[truple.KV[T, R]] {
	generater := stm.activate
	generaterR := stmR.activate
	return Stream[truple.KV[T, R]]{
		parallelism: stm.parallelism,
		activate: func(parallelism int) iterator.Iter[iterator.Iter[truple.KV[T, R]]] {

			segementer := generater(parallelism)
			segementerR := generaterR(parallelism)

			return func() (pr iterator.Iter[truple.KV[T, R]], ok bool) {
				pull, o1 := segementer()
				pullR, o2 := segementerR()
				if o1 && o2 {
					return func() (val truple.KV[T, R], ok bool) {
						t, ok := pull()
						r, ok2 := pullR()
						if ok && ok2 {
							return truple.KV[T, R]{
								Key: t,
								Val: r,
							}, ok
						}

						return
					}, true
				}

				return nil, false

			}
		},
	}
}
