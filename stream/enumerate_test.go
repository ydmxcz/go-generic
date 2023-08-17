package stream_test

import (
	"fmt"
	"testing"

	"github.com/ydmxcz/go-generic/iterator"
	"github.com/ydmxcz/go-generic/stream"
)

func TestEnumerate(t *testing.T) {
	stm := iterator.SplitableSliceOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	k := stream.CollectToSlice(stream.Enumerate(stream.New(stm)))
	fmt.Println(k)

}
