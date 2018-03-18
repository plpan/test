package main

import "testing"

func Benchmark_testType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testType()
	}
}

func Benchmark_testValue(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValue()
	}
}
