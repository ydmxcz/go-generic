package stream_test

import (
	"fmt"
	"testing"

	"github.com/ydmxcz/go-generic/iterator"
	"github.com/ydmxcz/go-generic/stream"
)

func TestMin(t *testing.T) {
	// sli := iterator.SplitableSliceOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sli := iterator.SplitableSliceOf(61, 4, 6, 8, 60, 12, 14, 16, 18, 62)
	allMatch := stream.Min(stream.New(sli))
	fmt.Println(allMatch)
}
