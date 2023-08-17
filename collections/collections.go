package collections

import (
	"github.com/ydmxcz/go-generic/iterator"
	"github.com/ydmxcz/go-generic/stream"
)

// type Segmenter[T any] func(parallelism int) iterator.Iter[iterator.Iter[T]]

type Iterable[T any] interface {
	Iter() iterator.Iter[T]

	SplitableIter() iterator.SplitableIter[T]
}

type Streamable[T any] interface {
	Stream(parallelism ...int) stream.Stream[T]
}

type Collection[T any] interface {
	Iterable[T]

	Add(elem T) bool

	AddAll(elems ...T) int

	// AddCollection(collection Collection[T]) int

	Delete(elem T) bool

	DeleteAll(elem ...T) int

	// DeleteCollection(collection Collection[T]) int

	Clear()

	Len() int

	Empty() bool

	ToSlice() []T
}

type List[T any] interface {
	Collection[T]

	Get(idx int)

	Index(elem T) int
}

type Queue[T any] interface {
	Push(val T) bool
	Pop() (val T, ok bool)
	Len() int
	Peek() (val T, ok bool)
}
