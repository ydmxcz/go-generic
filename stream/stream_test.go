package stream_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/ydmxcz/go-generic/stream"
)

func TestChannelStream(t *testing.T) {
	testCases := []struct {
		desc string

		ch          chan int
		parallelism int
	}{
		{
			desc:        "parallem-1(signle goroutine)",
			ch:          make(chan int, 10),
			parallelism: 1,
		}, {
			desc:        "parallem-4(more goroutine)",
			ch:          make(chan int, 10),
			parallelism: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(_ *testing.T) {
			go func() {
				time.Sleep(time.Second)
				fmt.Println("write start")
				for i := 0; i < 40; i++ {
					fmt.Println("write to channel:", i)
					tC.ch <- i
					// time.Sleep(time.Second >> 1)
				}
				close(tC.ch)
				fmt.Println("exit write")
			}()
			s := stream.Channel(tC.ch, tC.parallelism)
			stream.Collect(stream.Map(stream.Filter(s, func(n int) bool { return n%2 == 0 }), func(n int) string {
				return strconv.Itoa(n)
			}), func(s string) {
				fmt.Println("collect element", s)
			})
		})
	}
}
