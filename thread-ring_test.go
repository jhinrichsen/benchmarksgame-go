package main

import "testing"

func TestThreadRingE4(t *testing.T) {
	id := threadRing(1000)
	if id != 498 {
		t.Fatalf("want %v but got %v\n", 498, id)
	}
}

func TestThreadRingE6(t *testing.T) {
	id := threadRing(5e6)
	if id != 181 {
		t.Fatalf("want %v but got %v\n", 181, id)
	}
}

func TestThreadRingE7(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long test")
	}
	id := threadRing(5e7)
	if id != 292 {
		t.Fatalf("want %v but got %v\n", 292, id)
	}
}

func BenchmarkThreadRingE2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threadRing(5e2)
	}
}
func BenchmarkThreadRingE3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threadRing(5e3)
	}
}
func BenchmarkThreadRingE4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threadRing(5e4)
	}
}
func BenchmarkThreadRingE5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threadRing(5e5)
	}
}
func BenchmarkThreadRingE6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threadRing(5e6)
	}
}
