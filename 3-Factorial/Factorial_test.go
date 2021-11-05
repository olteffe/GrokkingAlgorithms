package main

import (
	"testing"
)

func Test_factorialRecursive(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "test n=0",
			args: args{
				n: 0,
			},
			want: 1,
		},
		{
			name: "test n=1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "test n=5",
			args: args{
				n: 5,
			},
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorialRecursive(tt.args.n); got != tt.want {
				t.Errorf("factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factorialIterative(t *testing.T) {
	type args struct {
		number uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "test n=0",
			args: args{
				number: 0,
			},
			want: 1,
		},
		{
			name: "test n=1",
			args: args{
				number: 1,
			},
			want: 1,
		},
		{
			name: "test n=5",
			args: args{
				number: 5,
			},
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorialIterative(tt.args.number); got != tt.want {
				t.Errorf("factorialIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFactorialRecursive(b *testing.B) {
	benchData := map[string]struct {
		size uint64
	}{
		"n=100": {size: 100},
		"n=10k": {size: 10000},
	}
	b.ResetTimer()
	for benchName, data := range benchData {
		b.Run(benchName, func(b *testing.B) {
			for idx := 0; idx < b.N; idx++ {
				factorialRecursive(data.size)
			}
		})
	}
}

func BenchmarkFactorialIterative(b *testing.B) {
	benchData := map[string]struct {
		size uint64
	}{
		"n=100": {size: 100},
		"n=10k": {size: 10000},
	}
	b.ResetTimer()
	for benchName, data := range benchData {
		b.Run(benchName, func(b *testing.B) {
			for idx := 0; idx < b.N; idx++ {
				factorialIterative(data.size)
			}
		})
	}
}
