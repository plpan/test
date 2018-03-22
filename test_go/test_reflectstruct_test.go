package main

import "testing"

func BenchmarkStructIter(b *testing.B) {
	b.ReportAllocs()
	u := User{1, "stupig", 2, "3"}
	for i := 0; i < b.N; i++ {
		StructIter(&u)
	}
}
