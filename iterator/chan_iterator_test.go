package iterator_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/ydmxcz/go-generic/iterator"
)

func TestChanIter(t *testing.T) {
	intChan := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("write to channel:", i+666)
			intChan <- (i + 666)
			time.Sleep(time.Second)
		}
		close(intChan)
		fmt.Println("exit write")
	}()
	iterator.Channel(intChan).All(func(i int) bool {
		fmt.Println("collect element:", i)
		return true
	})

}

func TestSplitableChannel(t *testing.T) {
	intChan := make(chan int, 10)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("write to channel:", i+666)
			intChan <- (i + 666)
			time.Sleep(time.Second >> 1)
		}
		close(intChan)
		fmt.Println("exit write")
	}()
	splitIter := iterator.SplitableChannel(intChan)
	goroutineNum := 4
	iter := splitIter(goroutineNum)
	wg := sync.WaitGroup{}
	wg.Add(goroutineNum)
	for i := 0; i < goroutineNum; i++ {
		i2, b := iter()
		if !b {
			t.Fatal("create iterator fail")
		}
		go func(num int, iter iterator.Iter[int]) {
			iter.All(func(i int) bool {
				fmt.Println("goroutine", num, " collect element:", i)
				return true
			})
			wg.Done()
		}(i, i2)
	}
	wg.Wait()
	// iterator.Channel(intChan).All(func(i int) bool {
	// 	fmt.Println("collect element:", i)
	// 	return true
	// })

}
