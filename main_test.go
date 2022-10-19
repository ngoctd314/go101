package main

import (
	_ "net/http/pprof"
	"testing"
)

func BenchmarkFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn()
	}
}

func BenchmarkFn1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn1()
	}
}
