package main

import "testing"

func TestNonDivisibleSubset(t *testing.T) {
	var k int32
	k = 3
	nonDivisibleSubset(k, input)
}

func BenchmarkNonDivisibleSubset(b *testing.B) {
	b.ReportAllocs()
	var k int32
	k = 3
	nonDivisibleSubset(k, input)
}

func BenchmarkGoroutineNonDivisibleSubset(b *testing.B) {
	b.ReportAllocs()
	var k int32
	k = 3
	goroutinesNonDivisibleSubset(k, input)
}
