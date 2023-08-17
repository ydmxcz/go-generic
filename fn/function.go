package fn

type Function[T, R any] func(T) R

func BeforeFunc[T, V, R any](before Function[V, T], curr Function[T, R]) Function[V, R] {
	return func(v V) R {
		return curr(before(v))
	}
}

func AfterFunc[T, V, R any](curr Function[T, R], after Function[R, V]) Function[T, V] {
	return func(t T) V {
		return after(curr(t))
	}
}

// Represents a function that accepts two arguments and produces a result.
// This is the two-arity specialization of Function.
type BinFunction[T, U, R any] func(T, U) R
