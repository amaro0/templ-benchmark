package main

import (
	"testing"
)

func BenchmarkHello(b *testing.B) {
	b.Run("engine=native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helloNative()
		}
	})
	b.Run("engine=templ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helloTempl()
		}
	})
}

func BenchmarkTable(b *testing.B) {
	b.Run("engine=native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tableNative()
		}
	})
	b.Run("engine=templ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tableTempl()
		}
	})
}

func BenchmarkLayout(b *testing.B) {
	b.Run("engine=native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			layoutNative()
		}
	})
	b.Run("engine=templ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			layoutTempl()
		}
	})
}
