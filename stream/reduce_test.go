package stream_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ydmxcz/go-generic/iterator"
	"github.com/ydmxcz/go-generic/stream"
)

func TestWorldCount(t *testing.T) {
	iter := iterator.SplitableSliceOf("hello world", "hello world mcz", "hello world", "hello world", "hello world")
	count := stream.ReduceWith(stream.Map(stream.New(iter),
		func(str string) []string {
			return strings.Split(str, " ")
		}), 0,
		func(count int, val []string) int {
			return count + len(val)
		},
		func(accum, val int) int {
			return accum + val
		})
	fmt.Println(count)
}
