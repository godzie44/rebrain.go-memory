package examples

import (
	"fmt"
	"testing"
)

type S struct {
	a, b, c int64
	d, e, f string
	g, h, i float64
}

func byCopy() S {
	return S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func byPointer() *S {
	return &S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func BenchmarkMemoryStack(b *testing.B) {
	var s S

	for i := 0; i < b.N; i++ {
		s = byCopy()
	}
	_ = fmt.Sprintf("%v", s.a)
}

func BenchmarkMemoryHeap(b *testing.B) {
	var s *S
	for i := 0; i < b.N; i++ {
		s = byPointer()
	}
	_ = fmt.Sprintf("%v", s.a)
}

func BenchmarkMemoryHeap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := byPointer()
		if s.a != 1 {
			panic("a!=1")
		}
	}
}
