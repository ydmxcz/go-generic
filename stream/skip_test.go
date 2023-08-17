package stream_test

import (
	"fmt"
	"testing"

	"github.com/ydmxcz/go-generic/iterator"
	"github.com/ydmxcz/go-generic/stream"
)

func TestSkip(t *testing.T) {
	iter := iterator.SplitableSliceOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	stm := stream.New(iter)
	stream.Collect(stream.Skip(stm, 3), func(a int) {
		fmt.Print(a, " ")
	})

}

func TestSkipWhile(t *testing.T) {
	iter := iterator.SplitableSliceOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	stm := stream.New(iter)
	arr := stream.CollectToSlice(stream.SkipWhile(stm, func(a int) bool {
		fmt.Print(a, " ")
		return true
	}))
	fmt.Println(arr)

}
