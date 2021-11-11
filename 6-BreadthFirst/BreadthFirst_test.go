package main

import "testing"

func Test_breadthSearch(t *testing.T) {
	type args struct {
		graph inputMap
		name  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				graph: inputMap{"you": {"alice", "bob", "claire"}, "bob": {"anuj", "peggy"}, "alice": {"peggy"},
					"claire": {"thom", "jonny"}, "anuj": {}, "peggy": {}, "thom": {}, "jonny": {}},
				name: "you",
			},
			want: true,
		},
		{
			name: "invalid",
			args: args{
				graph: inputMap{"you": {"alice", "bob", "claire"}, "bob": {"anuj", "peggy"}, "alice": {"peggy"},
					"claire": {"jonny"}, "anuj": {}, "peggy": {}, "jonny": {}},
				name: "you",
			},
			want: false,
		},
		{
			name: "null object",
			args: args{
				graph: inputMap{"": {""}},
				name:  "you",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breadthSearch(tt.args.graph, tt.args.name); got != tt.want {
				t.Errorf("breadthSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_personIsSeller(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				name: "thom",
			},
			want: true,
		},
		{
			name: "invalid",
			args: args{
				name: "thomas",
			},
			want: false,
		},
		{
			name: "null object",
			args: args{
				name: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := personIsSeller(tt.args.name); got != tt.want {
				t.Errorf("personIsSeller() = %v, want %v", got, tt.want)
			}
		})
	}
}
