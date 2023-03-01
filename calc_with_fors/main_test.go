package main

import (
	"math/rand"
	"testing"
)

func BenchmarkCalcular240K(b *testing.B) {
	ms := make([]M, 240_000)
	for i := range ms {
		ms[i] = M{
			Tipo:  rand.Intn(4),
			Valor: rand.Float64(),
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		calcular(ms)
	}
}

func BenchmarkCalcular1240K(b *testing.B) {
	ms := make([]M, 240_000)
	for i := range ms {
		ms[i] = M{
			Tipo:  rand.Intn(4),
			Valor: rand.Float64(),
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		calcular1(ms)
	}
}

func BenchmarkCalcular2_240K(b *testing.B) {
	ms := make([]M, 240_000)
	for i := range ms {
		ms[i] = M{
			Tipo:  rand.Intn(4),
			Valor: rand.Float64(),
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		calcular2(ms)
	}
}

func BenchmarkCalcular3_240K(b *testing.B) {
	ms := make([]M, 240_000)
	for i := range ms {
		ms[i] = M{
			Tipo:  rand.Intn(4),
			Valor: rand.Float64(),
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		calcular3(ms)
	}
}

func BenchmarkCalcular4_240K(b *testing.B) {
	ms := make([]M, 240_000)
	for i := range ms {
		ms[i] = M{
			Tipo:  rand.Intn(4),
			Valor: rand.Float64(),
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		calcular4(ms)
	}
}

func BenchmarkCalcular5_240K(b *testing.B) {
	ms := make([]M, 240_000)
	for i := range ms {
		ms[i] = M{
			Tipo:  rand.Intn(4),
			Valor: rand.Float64(),
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		calcular5(ms)
	}
}

func BenchmarkCalcular6_240K(b *testing.B) {
	ms := make([]M, 240_000)
	for i := range ms {
		ms[i] = M{
			Tipo:  rand.Intn(4),
			Valor: rand.Float64(),
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		calcular6(ms)
	}
}

func BenchmarkCalcular7_240K(b *testing.B) {
	ms := make([]M, 240_000)
	for i := range ms {
		ms[i] = M{
			Tipo:  rand.Intn(4),
			Valor: rand.Float64(),
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		calcular7(ms)
	}
}
