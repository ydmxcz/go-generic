package iterator

import (
	"sync/atomic"
)

func Channel[T any](c chan T) Iter[T] {
	return func() (val T, ok bool) {
		val, ok = <-c
		return val, ok
	}
}

func SplitableChannel[T any](c chan T) func(int) Iter[Iter[T]] {
	return func(parallelism int) Iter[Iter[T]] {
		var idx int32 = 0
		if parallelism <= 0 {
			return func() (Iter[T], bool) {
				if atomic.LoadInt32(&idx) == 0 {
					atomic.AddInt32(&idx, 1)
				}
				return Channel[T](c), true
			}
		}
		return func() (pull Iter[T], ok bool) {
			if atomic.LoadInt32(&idx) >= int32(parallelism) {
				return nil, false
			}
			atomic.AddInt32(&idx, 1)
			return Channel[T](c), true
		}
	}
}
