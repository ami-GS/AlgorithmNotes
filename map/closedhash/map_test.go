package closedhash

import (
	"math/rand"
	"testing"
)

func BenchmarkInsert(b *testing.B) {
	m := NewMap(10000019)

	for i := 0; i < b.N/100; i++ {
		key := rand.Int()
		val := rand.Int()
		m.Insert(key, val)
	}
}
