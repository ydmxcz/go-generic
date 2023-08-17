package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ydmxcz/go-generic/iterator"
	"github.com/ydmxcz/go-generic/stream"
)

func TestSum(t *testing.T) {
	sli := iterator.SplitableSliceOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := stream.Sum(stream.New(sli))
	assert.Equal(t, 220, sum)
}

func TestSum_Parallel(t *testing.T) {
	sli := iterator.SplitableSliceOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := stream.Sum(stream.New(sli, 4))
	assert.Equal(t, 220, sum)

}
