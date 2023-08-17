package fn_test

import (
	"fmt"
	"testing"

	"github.com/ydmxcz/go-generic/fn"
)

type testObj struct {
	ID int
}

func TestChain(t *testing.T) {
	v := fn.Chain[*testObj](func(elem *testObj) error {
		fmt.Println(elem)
		elem.ID = 666
		return nil
	})
	v = v.Around(func(elem *testObj, next fn.Chain[*testObj]) error {
		fmt.Println("1----AAA", elem.ID)
		next(elem)
		fmt.Println("1----BBB", elem.ID)
		return nil
	})
	v = v.Around(func(elem *testObj, next fn.Chain[*testObj]) error {
		fmt.Println("2----aaa", elem.ID)
		next(elem)
		fmt.Println("2----bbb", elem.ID)
		return nil
	})
	obj := &testObj{ID: 10}
	v(obj)
	// fn.Around(func(c fn.Chain[*testObj]) {
	// 	fmt.Println("before-1")
	// 	c(nil)
	// 	fmt.Println("after-1")
	// })
}
