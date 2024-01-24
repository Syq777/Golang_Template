package Chapter01

import (
	"reflect"
	"testing"
)

func TestDynamicSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Test Case 1",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 3, 6, 10, 15},
		},
		{
			name: "Test Case 2",
			nums: []int{0, 0, 0, 0, 0},
			want: []int{0, 0, 0, 0, 0},
		},
		{
			name: "Test Case 3",
			nums: []int{5, 4, 3, 2, 1},
			want: []int{1, 5, 13, 26, 52},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DanamicSum(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
