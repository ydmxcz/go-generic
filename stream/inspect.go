package stream

import "github.com/ydmxcz/go-generic/iterator"

func Inspect[T any](stm Stream[T], inspect func(val T)) Stream[T] {
	generater := stm.activate

	return Stream[T]{
		parallelism: stm.parallelism,
		activate: func(parallelism int) iterator.Iter[iterator.Iter[T]] {

			segementer := generater(parallelism)

			return func() (pr iterator.Iter[T], ok bool) {
				if pull, o1 := segementer(); o1 {
					return func() (val T, ok bool) {
						if val, ok = pull(); ok {
							inspect(val)
							return val, ok
						}
						return
					}, true
				}
				return nil, false

			}
		},
	}
}
