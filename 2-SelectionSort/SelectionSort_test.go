package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_findSmallest(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1item",
			args: args{
				arr: []int{1},
			},
			want: 0,
		},
		{
			name: "test10item",
			args: args{
				arr: []int{9, -8, 7, 6, 5, 4, -3, 2, 1, 0},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSmallest(tt.args.arr); got != tt.want {
				t.Errorf("findSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name       string
		args       args
		wantNewArr []int
	}{
		{
			name: "test0item",
			args: args{
				arr: []int{},
			},
			wantNewArr: []int{},
		},
		{
			name: "test1item",
			args: args{
				arr: []int{1},
			},
			wantNewArr: []int{1},
		},
		{
			name: "test10item",
			args: args{
				arr: []int{-9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			},
			wantNewArr: []int{-9, 0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewArr := selectionSort(tt.args.arr); !reflect.DeepEqual(gotNewArr, tt.wantNewArr) {
				t.Errorf("selectionSort() = %v, want %v", gotNewArr, tt.wantNewArr)
			}
		})
	}
}

func Test_goSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test0item",
			args: args{
				arr: []int{},
			},
			want: []int{},
		},
		{
			name: "test1item",
			args: args{
				arr: []int{1},
			},
			want: []int{1},
		},
		{
			name: "test10item",
			args: args{
				arr: []int{-9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			},
			want: []int{-9, 0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
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
				selectionSort(randomSlice)
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
