package sum

import "testing"

func TestInts(t *testing.T) {
	tt := []struct {
		numbers []int
		sum     int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{nil, 0},
		{[]int{1, -1}, 0},
	}

	for _, tc := range tt {
		s := Ints(tc.numbers...)
		if s != tc.sum {
			t.Errorf("sum of %v should be %v; got %v", tc.numbers, tc.sum, s)
		}
	}
}
