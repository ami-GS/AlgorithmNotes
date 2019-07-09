package openhash

import (
	"math/rand"
	"testing"
)

func BenchmarkInsert(b *testing.B) {
	m := NewMap(10000019)

	for i := 0; i < b.N; i++ {
		key := rand.Int()
		val := rand.Int()
		//fmt.Println(key, val)
		m.Insert(key, val)
	}
}
