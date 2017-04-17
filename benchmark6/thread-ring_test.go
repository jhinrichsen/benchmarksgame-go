package benchmark6

import "testing"

func BenchmarkThreadRing6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threadRing6(50000000)
	}
}
