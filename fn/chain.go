package fn

type Chain[T any] func(T) error

func (c Chain[T]) Around(around func(elem T, next Chain[T]) error) Chain[T] {
	return func(t T) error {
		return around(t, c)
	}
}
