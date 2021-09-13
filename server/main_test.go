package main

import (
	"testing"
)

// go test . -bench=. -benchmem
func BenchmarkMakeUsers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeUsers(1000)
	}
}