package fn

type Predicate[T any] func(T) bool

func (p Predicate[T]) And(other Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return p(t) && other(t)
	}
}

func (p Predicate[T]) Negate() Predicate[T] {
	return func(t T) bool {
		return !p(t)
	}
}

func (p Predicate[T]) Or(other Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return p(t) || other(t)
	}
}

func Not[T any](target Predicate[T]) Predicate[T] {
	return target.Negate()
}

type BinPredicate[T, U any] func(T, U) bool

func (curr BinPredicate[T, U]) And(other BinPredicate[T, U]) BinPredicate[T, U] {
	return func(t T, u U) bool {
		return curr(t, u) && other(t, u)
	}
}

func (curr BinPredicate[T, U]) Or(other BinPredicate[T, U]) BinPredicate[T, U] {
	return func(t T, u U) bool {
		return curr(t, u) && other(t, u)
	}
}

func (curr BinPredicate[T, U]) negate() BinPredicate[T, U] {
	return func(t T, u U) bool {
		return !curr(t, u)
	}
}
