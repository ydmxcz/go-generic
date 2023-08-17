package linkedlist_test

import (
	"fmt"
	"testing"

	"github.com/ydmxcz/go-generic/collections/linkedlist"
)

func TestLinkedList(t *testing.T) {
	l := linkedlist.New[int]()
	e := l.PushBack(1)
	l.Remove(e)
	// for i := 0; i < 10; i++ {
	// 	l.PushBack(i)
	// }
	// iter, stop := iterator.CastToPull(l.Iter().All)
	// for v, ok := iter(); ok; v, ok = iter() {
	// 	fmt.Println(v)
	// }
	// stop()
	fmt.Println(l.ToSlice())
}

func TestLinkedListSplitableIter(t *testing.T) {
	l := linkedlist.New[int]()
	for i := 1; i <= 10; i++ {
		l.PushBack(i)
	}
	iter := l.SplitableIter()(3)
	for iter1, ok := iter(); ok; iter1, ok = iter() {
		for v1, o1 := iter1(); o1; v1, o1 = iter1() {
			fmt.Println(v1)
		}
		fmt.Println("====================")
	}
}

func TestBoardCast(t *testing.T) {
	// map[string]strings.Bui
	l := linkedlist.New[int]()
	fmt.Println(l.ToSlice())
	l.PushBack(666)
	fmt.Println(l.ToSlice())
	for i := 1; i <= 10; i++ {
		l.PushBack(i)
	}
	fmt.Println(l.ToSlice())
}
