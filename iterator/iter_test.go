package iterator

import (
	"fmt"
	"testing"
)

func TestMapIterator(t *testing.T) {
	m := map[int]int{
		1: 111,
		2: 222,
		3: 333,
		4: 444,
		5: 555,
		6: 666,
		7: 777,
	}
	iter := Map(m)
	for kv, ok := iter(); ok; kv, ok = iter() {
		fmt.Println(kv)
	}

	fmt.Println("===============")
	iter = Map(m)
	for kv, ok := iter(); ok; kv, ok = iter() {
		fmt.Println(kv)
	}
}
