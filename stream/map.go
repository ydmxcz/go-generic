package stream

import "github.com/ydmxcz/go-generic/iterator"

func Map[T, R any](stm Stream[T], mapF func(T) R) Stream[R] {
	generater := stm.activate

	return Stream[R]{
		parallelism: stm.parallelism,
		activate: func(parallelism int) iterator.Iter[iterator.Iter[R]] {

			segementer := generater(parallelism)

			return func() (pr iterator.Iter[R], ok bool) {
				if pull, o1 := segementer(); o1 {
					return func() (val R, ok bool) {
						if t, ok := pull(); ok {
							return mapF(t), ok
						}
						return
					}, true
				}
				return nil, false

			}
		},
	}
}
