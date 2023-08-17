package stream

import (
	"github.com/ydmxcz/go-generic/fn"
	"github.com/ydmxcz/go-generic/iterator"
)

func CollectToSlice[T any](stm Stream[T]) []T {
	arr := make([]T, 4)
	Collect(stm, func(t T) { arr = append(arr, t) })
	return arr
}

func Collect[T any](stm Stream[T], collectTo fn.Consumer[T]) {
	if stm.parallelism <= 1 {
		iterGenerators := stm.activate(stm.parallelism)
		pull, b := iterGenerators()
		if !b {
			return
		}
		for val, ok := pull(); ok; val, ok = pull() {
			collectTo(val)
		}
	} else {

		doParallel(stm.activate(stm.parallelism),
			func(_ int, pull iterator.Iter[T]) {
				// log.Println("goroutine", idx, "start")
				for val, ok := pull(); ok; val, ok = pull() {
					collectTo(val)
				}
				// log.Println("goroutine", idx, "end")
			}).Wait()
	}
}
