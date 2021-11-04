package main

import "testing"

// makeRange generate slice [0, 1, 2, .. max)
func makeRange(max int) []int {
	slice := make([]int, max)
	for i := range slice {
		slice[i] = i
	}
	return slice
}

var (
	const100  = makeRange(100)
	const100k = makeRange(100000)
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		list []int
		item int
	}
	tests := []struct {
		name    string
		args    args
		wantPos int
	}{
		{
			name: "test1item",
			args: args{
				list: makeRange(1),
				item: 0,
			},
			wantPos: 0,
		},
		{
			name: "test10item",
			args: args{
				list: makeRange(10),
				item: 5,
			},
			wantPos: 5,
		},
		{
			name: "test1000item",
			args: args{
				list: makeRange(1000),
				item: 999,
			},
			wantPos: 999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPos := BinarySearch(tt.args.list, tt.args.item); gotPos != tt.wantPos {
				t.Errorf("binarySearch() = %v, want %v", gotPos, tt.wantPos)
			}
		})
	}
}

func BenchmarkBinarySearch100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch(const100, 99)
	}
}

func BenchmarkBinarySearch100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch(const100k, 99999)
	}
}

func TestSerialSearch(t *testing.T) {
	type args struct {
		list []int
		item int
	}
	tests := []struct {
		name    string
		args    args
		wantPos int
	}{
		{
			name: "test1item",
			args: args{
				list: makeRange(1),
				item: 0,
			},
			wantPos: 0,
		},
		{
			name: "test10item",
			args: args{
				list: makeRange(10),
				item: 5,
			},
			wantPos: 5,
		},
		{
			name: "test1000item",
			args: args{
				list: makeRange(1000),
				item: 999,
			},
			wantPos: 999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPos := SerialSearch(tt.args.list, tt.args.item); gotPos != tt.wantPos {
				t.Errorf("SerialSearch() = %v, want %v", gotPos, tt.wantPos)
			}
		})
	}
}

func BenchmarkSerialSearch100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SerialSearch(const100, 99)
	}
}

func BenchmarkSerialSearch100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SerialSearch(const100k, 99999)
	}

}
