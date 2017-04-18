package main

import "testing"

func TestThreadRingE4(t *testing.T) {
	id := threadRing(1000, 0)
	if id != 498 {
		t.Fatalf("want %v but got %v\n", 498, id)
	}
}

func TestThreadRingE6(t *testing.T) {
	id := threadRing(5e6, 0)
	if id != 181 {
		t.Fatalf("want %v but got %v\n", 181, id)
	}
}

func TestThreadRingE7(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long test")
	}
	id := threadRing(5e7, 0)
	if id != 292 {
		t.Fatalf("want %v but got %v\n", 292, id)
	}
}

func btr(b *testing.B, n, size int) {
	for i := 0; i < b.N; i++ {
		threadRing(5e2, size)
	}
}

func BenchmarkThreadRingBufferedE2(b *testing.B) {
	btr(b, 5e2, size)
}
func BenchmarkThreadRingBufferedE3(b *testing.B) {
	btr(b, 5e3, size)
}
func BenchmarkThreadRingBufferedE4(b *testing.B) {
	btr(b, 5e4, size)
}
func BenchmarkThreadRingBufferedE5(b *testing.B) {
	btr(b, 5e5, size)
}
func BenchmarkThreadRingBufferedE6(b *testing.B) {
	btr(b, 5e6, size)
}
func BenchmarkThreadRingUnbufferedE2(b *testing.B) {
	btr(b, 5e2, 0)
}
func BenchmarkThreadRingUnbufferedE3(b *testing.B) {
	btr(b, 5e3, 0)
}
func BenchmarkThreadRingUnbufferedE4(b *testing.B) {
	btr(b, 5e4, 0)
}
func BenchmarkThreadRingUnbufferedE5(b *testing.B) {
	btr(b, 5e5, 0)
}
func BenchmarkThreadRingUnbufferedE6(b *testing.B) {
	btr(b, 5e6, 0)
}
