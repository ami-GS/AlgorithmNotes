package cuckoo

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkInsert(b *testing.B) {
	m := NewMap(10000019)
	fmt.Println(b.N)

	for i := 0; i < b.N; i++ {
		key := rand.Int()
		val := rand.Int()
		m.Insert(key, val)
	}
}
