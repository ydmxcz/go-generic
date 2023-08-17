package stream_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ydmxcz/go-generic/iterator"
	"github.com/ydmxcz/go-generic/stream"
)

func TestChain(t *testing.T) {
	s := stream.New(iterator.SplitableSliceOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), 4)
	s2 := stream.New(iterator.SplitableSliceOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), 4)

	stream.Collect(stream.Chain(s, s2), func(a int) {
		fmt.Println(a)
	})
}

func TestXxx(t *testing.T) {
	_, cancel := context.WithCancel(context.Background())
	cancel()
	cancel()

}
