package fn

type Supplier[T any] func() T

type BinOperator[T any] BinFunction[T, T, T]
