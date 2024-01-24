package Chapter02

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want []int
	}{
		{
			name: "Test case 1",
			arr:  []int{1, 2, 3, 4, 5},
			k:    3,
			want: []int{3, 4, 5, 1, 2},
		},
		{
			name: "Test case 2",
			arr:  []int{7, 8, 9, 10, 11},
			k:    5,
			want: []int{8, 9, 10, 11, 7},
		},
		{
			name: "Test case 3",
			arr:  []int{2, 3, 4, 5, 6},
			k:    1,
			want: []int{3, 4, 5, 6, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateArray(tt.arr, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
