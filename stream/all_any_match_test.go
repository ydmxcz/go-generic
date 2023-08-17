package stream_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ydmxcz/go-generic/iterator"
	"github.com/ydmxcz/go-generic/stream"
)

func MyTest[T any](t *testing.T, testData []T) {

}

func TestAllMatch(t *testing.T) {
	testCases := []struct {
		desc        string
		iter        iterator.SplitableIter[int]
		parallelism int
		expect      bool
	}{
		{
			desc:        "[2, 4, 6, 8, 10, 12, 14, 16, 18, 21]---parallelism:1",
			iter:        iterator.SplitableSliceOf(2, 4, 6, 8, 10, 12, 14, 16, 18, 21),
			parallelism: 1,
			expect:      false,
		},
		{
			desc:        "[2, 4, 6, 8, 10, 12, 14, 16, 18, 21]---parallelism:4",
			iter:        iterator.SplitableSliceOf(2, 4, 6, 8, 10, 12, 14, 16, 18, 21),
			parallelism: 4,
			expect:      false,
		}, {
			desc:        "[2, 4, 6, 8, 10, 12, 14, 16, 18]---parallelism:1",
			iter:        iterator.SplitableSliceOf(2, 4, 6, 8, 10, 12, 14, 16, 18),
			parallelism: 1,
			expect:      true,
		}, {
			desc:        "[2, 4, 6, 8, 10, 12, 14, 16, 18]---parallelism:4",
			iter:        iterator.SplitableSliceOf(2, 4, 6, 8, 10, 12, 14, 16, 18),
			parallelism: 4,
			expect:      true,
		}, {
			desc:        "[1, 2, 4, 6, 8, 10, 12, 14, 16, 18]---parallelism:1",
			iter:        iterator.SplitableSliceOf(1, 2, 4, 6, 8, 10, 12, 14, 16, 18),
			parallelism: 1,
			expect:      false,
		}, {
			desc:        "[1, 2, 4, 6, 8, 10, 12, 14, 16, 18]---parallelism:4",
			iter:        iterator.SplitableSliceOf(1, 2, 4, 6, 8, 10, 12, 14, 16, 18),
			parallelism: 4,
			expect:      false,
		}, {
			desc:        "[2, 4, 6, 8, 10, 11, 12, 14, 16, 18]---parallelism:1",
			iter:        iterator.SplitableSliceOf(1, 2, 4, 6, 8, 10, 11, 12, 14, 16, 18),
			parallelism: 1,
			expect:      false,
		}, {
			desc:        "[2, 4, 6, 8, 10, 11, 12, 14, 16, 18]---parallelism:4",
			iter:        iterator.SplitableSliceOf(1, 2, 4, 6, 8, 10, 11, 12, 14, 16, 18),
			parallelism: 4,
			expect:      false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			allMatch := stream.AllMatch(stream.New(tC.iter, tC.parallelism), func(n int) bool {
				return n%2 == 0
			})
			assert.Equal(t, tC.expect, allMatch)
		})
	}
}

func TestAnyMatch(t *testing.T) {
	sli := iterator.SplitableSliceOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// sli := iterator.SplitableSliceOf(1, 3, 5, 7, 9, 11, 13, 15, 17, 19)
	// sli := iterator.SplitableSliceOf(2, 4, 6, 8, 10, 12, 14, 16, 18, 21)

	allMatch := stream.AnyMatch(stream.New(sli), func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(allMatch)
}
