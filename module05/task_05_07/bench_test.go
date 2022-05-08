package main

import "testing"

/* func BenchmarkPutInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			b.StopTimer()
			PutInt(j)
			b.StartTimer()
		}
	}
} */

func BenchmarkSetInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			b.StopTimer()
			SetInt(j)
			b.StartTimer()
		}
	}
}
