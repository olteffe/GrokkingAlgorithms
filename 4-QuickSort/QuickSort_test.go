package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_quickSortSlice(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test0item",
			args: args{
				array: []int{},
			},
			want: []int{},
		},
		{
			name: "test1item",
			args: args{
				array: []int{1},
			},
			want: []int{1},
		},
		{
			name: "test10item",
			args: args{
				array: []int{-9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			},
			want: []int{-9, 0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSortSlice(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSortStruct(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test0item",
			args: args{
				array: []int{},
			},
			want: []int{},
		},
		{
			name: "test1item",
			args: args{
				array: []int{1},
			},
			want: []int{1},
		},
		{
			name: "test10item",
			args: args{
				array: []int{-9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			},
			want: []int{-9, 0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSortStruct(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSortStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_goSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test0item",
			args: args{
				array: []int{},
			},
			want: []int{},
		},
		{
			name: "test1item",
			args: args{
				array: []int{1},
			},
			want: []int{1},
		},
		{
			name: "test10item",
			args: args{
				array: []int{-9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			},
			want: []int{-9, 0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkQuickSortStruct(b *testing.B) {
	benchData := map[string]struct {
		size int
	}{
		"size=100":  {size: 100},
		"size=100k": {size: 100000},
	}
	b.ResetTimer()
	for benchName, data := range benchData {
		b.StopTimer()
		randomSlice := rand.Perm(data.size)
		b.StartTimer()
		b.Run(benchName, func(b *testing.B) {
			for idx := 0; idx < b.N; idx++ {
				quickSortStruct(randomSlice)
			}
		})
	}
}

func BenchmarkQuickSortSlice(b *testing.B) {
	benchData := map[string]struct {
		size int
	}{
		"size=100":  {size: 100},
		"size=100k": {size: 100000},
	}
	b.ResetTimer()
	for benchName, data := range benchData {
		b.StopTimer()
		randomSlice := rand.Perm(data.size)
		b.StartTimer()
		b.Run(benchName, func(b *testing.B) {
			for idx := 0; idx < b.N; idx++ {
				quickSortSlice(randomSlice)
			}
		})
	}
}

func BenchmarkGoSort(b *testing.B) {
	benchData := map[string]struct {
		size int
	}{
		"size=100":  {size: 100},
		"size=100k": {size: 100000},
	}
	b.ResetTimer()
	for benchName, data := range benchData {
		b.StopTimer()
		randomSlice := rand.Perm(data.size)
		b.StartTimer()
		b.Run(benchName, func(b *testing.B) {
			for idx := 0; idx < b.N; idx++ {
				goSort(randomSlice)
			}
		})
	}
}
