package main

import "testing"

/*
func TestThreadRingShort1(t *testing.T) {
	id := threadRing(1000)
	if id != 498 {
		t.Fatalf("expecting %v but got %v\n", 498, id)
	}
}

func TestThreadRingShort2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long test")
	}
	id := threadRing(5e6)
	if id != 181 {
		t.Fatalf("expecting %v but got %v\n", 181, id)
	}
}

func TestThreadRingLong(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long test")
	}
	id := threadRing(5e7)
	if id != 292 {
		t.Fatalf("expecting %v but got %v\n", 292, id)
	}
}
*/
func BenchmarkThreadRing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threadRing(5e3)
	}
}
