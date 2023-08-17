package iterator

func SliceOf[T any](s ...T) Iter[T] {
	return Slice(s)
}

func Slice[T any](s []T) Iter[T] {
	var idx = 0
	return func() (val T, ok bool) {
		ok = idx < len(s)
		if ok {
			val = s[idx]
			idx++
		}
		return
	}
}

func SplitableSliceOf[T any](s ...T) SplitableIter[T] {
	return SplitableSlice(s)
}

func SplitableSlice[T any](s []T) SplitableIter[T] {

	return func(parallelism int) Iter[Iter[T]] {
		idx := 0
		var step int
		if parallelism == 0 {
			step = len(s)
		} else {
			step = len(s) / parallelism
		}
		if parallelism <= 0 {
			return func() (Iter[T], bool) {
				if idx == 0 {
					idx++
					return Slice[T](s), true
				}
				return nil, false
			}
		}

		return func() (pull Iter[T], ok bool) {
			if idx >= len(s) {
				return nil, false
			}
			i := idx
			idx += step
			if i+step >= len(s) {
				return Slice[T](s[i:]), true
			}
			return Slice[T](s[i : i+step]), true
		}
	}
}
