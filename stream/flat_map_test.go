package stream_test

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/ydmxcz/go-generic/iterator"
	"github.com/ydmxcz/go-generic/stream"
)

func TestFlatMap(t *testing.T) {
	// sli := iterator.SplitableSliceOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sli := iterator.SplitableSliceOf("11111", "22222", "33333", "44444", "55555", "66666", "77777", "88888", "99999")
	var count int64
	stream.Collect(stream.Inspect(
		stream.Parallel(stream.FlatMap(stream.New(sli),
			func(s string) stream.Stream[byte] {
				return stream.New(iterator.SplitableSliceOf([]byte(s)...))
			}), 4),
		func(b byte) {
			fmt.Println(string(b))
		}),
		func(_ byte) {
			atomic.AddInt64(&count, 1)
		})
	fmt.Println("count:", count)
}
