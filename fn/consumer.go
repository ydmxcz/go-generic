package fn

type Consumer[T any] func(T)

func (c Consumer[T]) Around(before, after Consumer[T]) Consumer[T] {
	return func(t T) {
		before(t)
		c(t)
		after(t)
	}
}

func (c Consumer[T]) Before(before Consumer[T]) Consumer[T] {
	return func(t T) {
		before(t)
		c(t)
	}
}

func (c Consumer[T]) After(after Consumer[T]) Consumer[T] {
	return func(t T) {
		c(t)
		after(t)
	}
}

// Represents an operation that accepts two input arguments and returns no result.
// This is the two-arity specialization of Consumer.
// Unlike most other functional interfaces,
// BiConsumer is expected to operate via side-effects.
type BinConsumer[T, U any] func(T, U)

func (curr BinConsumer[T, U]) Around(before, after BinConsumer[T, U]) BinConsumer[T, U] {
	return func(t T, u U) {
		before(t, u)
		curr(t, u)
		after(t, u)
	}
}

func (curr BinConsumer[T, U]) Before(before BinConsumer[T, U]) BinConsumer[T, U] {
	return func(t T, u U) {
		before(t, u)
		curr(t, u)
	}
}

func (curr BinConsumer[T, U]) After(after BinConsumer[T, U]) BinConsumer[T, U] {
	return func(t T, u U) {
		curr(t, u)
		after(t, u)
	}
}
