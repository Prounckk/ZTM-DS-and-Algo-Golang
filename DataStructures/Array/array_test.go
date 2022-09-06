package array

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		values []interface{}
	}

	tests := []struct {
		name string
		args args
		want *MyArray
	}{
		{
			name: "empty NewArray",
			args: args{},
			want: &MyArray{},
		},
		{
			name: "NewArray",
			args: args{values: []interface{}{1, 2, 3}},
			want: &MyArray{data: []interface{}{1, 2, 3}, size: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
