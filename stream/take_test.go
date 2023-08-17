package stream_test

import (
	"fmt"
	"testing"

	"github.com/ydmxcz/go-generic/iterator"
	"github.com/ydmxcz/go-generic/stream"
)

func TestTake(t *testing.T) {
	sso := iterator.SplitableSliceOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	stm := stream.New(sso)

	arr := stream.CollectToSlice(stream.Take(stm, 5))
	fmt.Println(arr)
}
