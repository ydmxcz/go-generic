package fn

type Pull[T any] func() (T, bool)

type Push[T any] func(Predicate[T])

type PushPred[T any] Predicate[Predicate[T]]
