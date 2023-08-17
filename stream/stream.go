package stream

import (
	"sync"

	"github.com/ydmxcz/go-generic/fn"
	"github.com/ydmxcz/go-generic/iterator"
)

// type IterGenerater[T any] func(parallelism int) iterator.Iter[iterator.Iter[T]]

type Stream[T any] struct {
	activate    iterator.SplitableIter[T]
	parallelism int
}

func New[T any](generater iterator.SplitableIter[T], parallelism ...int) Stream[T] {
	p := 0
	if len(parallelism) != 0 {
		p = parallelism[0]
	}
	return Stream[T]{activate: generater, parallelism: p}
}

func Slice[T any](s []T, parallelism ...int) Stream[T] {
	return New(iterator.SplitableSlice(s), parallelism...)
}

func SliceOf[T any](parallelism int, s ...T) Stream[T] {
	return New(iterator.SplitableSlice(s), parallelism)
}

func Channel[T any](s chan T, parallelism ...int) Stream[T] {
	return New(iterator.SplitableChannel(s), parallelism...)
}

func doParallel[T any](iterGenerater iterator.Iter[iterator.Iter[T]],
	yield fn.BinConsumer[int, iterator.Iter[T]]) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	idx := 0
	for {
		iter, ok := iterGenerater()
		if !ok {
			break
		}
		wg.Add(1)
		go func(idx int, wg *sync.WaitGroup, iter iterator.Iter[T]) {
			yield(idx, iter)
			wg.Done()
		}(idx, wg, iter)
		idx++
	}
	return wg
}

func Parallel[T any](stm Stream[T], parallelism int) Stream[T] {
	stm.parallelism = parallelism
	return stm
}
