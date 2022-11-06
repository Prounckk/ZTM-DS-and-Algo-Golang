package threes

import (
	"testing"
)

func TestBinaryTree_search(t *testing.T) {
	type fields struct {
		root *Node
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "empty",
			fields: fields{root: &Node{}},
			args:   args{5},
			want:   false,
		},
		{name: "result in the bottom ",
			fields: fields{root: &Node{data: 6, left: &Node{data: 4, left: nil, right: nil}, right: &Node{data: 10, left: nil, right: nil}}},
			args:   args{10},
			want:   true,
		},
		{name: "result in the root ",
			fields: fields{root: &Node{data: 6, left: nil, right: nil}},
			args:   args{6},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &BinaryTree{
				root: tt.fields.root,
			}
			if got := tr.search(tt.args.data); got != tt.want {
				t.Errorf("BinaryTree.search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_insert(t *testing.T) {
	type args struct {
		data int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty",
			args: args{},
			want: false,
		},
		{name: "no results",
			args: args{88},
			want: false,
		},
		{name: "found",
			args: args{33},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &BinaryTree{}
			tr.insert(6)
			tr.insert(5)
			tr.insert(10)
			tr.insert(13)
			tr.insert(33)
			tr.insert(4)
			if got := tr.search(tt.args.data); got != tt.want {
				t.Errorf("BinaryTree.search() = %v, want %v", got, tt.want)
			}
		})
	}
}
