package stream_test

import (
	"fmt"
	"testing"

	"github.com/ydmxcz/go-generic/iterator"
	"github.com/ydmxcz/go-generic/stream"
)

func TestZip(t *testing.T) {
	sso := iterator.SplitableSliceOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sso2 := iterator.SplitableSliceOf(111, 222, 333, 444, 555, 666, 777, 888, 999)

	k := stream.CollectToSlice(stream.Zip(stream.New(sso), stream.New(sso2)))
	fmt.Println(k)

}
