package stream_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ydmxcz/go-generic/iterator"
	"github.com/ydmxcz/go-generic/stream"
)

func TestFunc(t *testing.T) {
	sli := iterator.SplitableSliceOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1666, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1888, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1999, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	s := stream.New(sli)

	stream.Collect(stream.Filter(s, func(a int) bool {
		time.Sleep(300 * time.Millisecond)
		return a%2 == 0
	}), func(a int) {
		fmt.Println(a)
	})
}
