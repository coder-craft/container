package sliceint

import "testing"

func TestAdd(t *testing.T) {
	var tests = []struct {
		source []int
		input  int
		want   []int
	}{
		{[]int{}, 5, []int{5}},
		{[]int{1}, 5, []int{1, 5}},
		{[]int{5}, 5, []int{5, 5}},
		{[]int{5, 5}, 5, []int{5, 5, 5}},
		{[]int{1, 1, 1, 1}, 5, []int{1, 1, 1, 1, 5}},
	}
	for _, test := range tests {
		si := NewSliceint()
		if len(test.source) > 0 {
			si = test.source
		}
		si.Add(test.input)
		count := 0
		for k, v := range test.want {
			if v == si[k] {
				count++
			} else {
				t.Errorf("[%v] not in source slice.", v)
			}
		}
		if count != len(test.want) {
			t.Errorf("Target slice length error.")
		}
	}
}
func TestInit(t *testing.T) {
	var tests = []struct {
		source []int
		want   []int
	}{
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
	}
	for _, test := range tests {
		si := NewSliceint(test.source[0], test.source[1], test.source[2], test.source[3])
		count := 0
		for k, v := range test.want {
			if v == si[k] {
				count++
			} else {
				t.Errorf("[%v] not in source slice.", v)
			}
		}
		if count != len(test.want) {
			t.Errorf("Target slice length error.")
		}
	}
}
func TestDelete(t *testing.T) {
	var tests = []struct {
		source []int
		input  int
		want   []int
	}{
		{[]int{}, 5, []int{}},
		{[]int{1}, 5, []int{1}},
		{[]int{5}, 5, []int{}},
		{[]int{1, 5}, 5, []int{1}},
		{[]int{5, 1}, 5, []int{1}},
		{[]int{1, 1, 5, 1}, 5, []int{1, 1, 1}},
		{[]int{1, 1, 1, 1}, 5, []int{1, 1, 1, 1}},
	}
	for index, test := range tests {
		si := NewSliceint()
		if len(test.source) > 0 {
			si = test.source
		}
		si.Delete(test.input)
		count := 0
		for k, v := range test.want {
			if v == si[k] {
				count++
			} else {
				t.Errorf("Test data:[%v],[%v] not in source slice.", index, v)
			}
		}
		if count != len(test.want) {
			t.Errorf("Test data:[%v],Target slice length error.", index)
		}
	}
	for index, test := range tests {
		si := NewSliceint()
		if len(test.source) > 0 {
			si = test.source
		}
		si.Delete(test.input, test.input, test.input)
		count := 0
		for k, v := range test.want {
			if v == si[k] {
				count++
			} else {
				t.Errorf("Test data:[%v],[%v] not in source slice.", index, v)
			}
		}
		if count != len(test.want) {
			t.Errorf("Test data:[%v],Target slice length error.", index)
		}
	}
}
func TestCheck(t *testing.T) {
	var tests = []struct {
		source []int
		input  int
		result bool
	}{
		{[]int{}, 5, false},
		{[]int{1}, 5, false},
		{[]int{5}, 5, true},
		{[]int{1, 5}, 5, true},
		{[]int{5, 1}, 5, true},
		{[]int{1, 1, 5, 1}, 5, true},
		{[]int{1, 1, 1, 1}, 5, false},
	}
	for index, test := range tests {
		si := NewSliceint()
		if len(test.source) > 0 {
			si = test.source
		}
		if si.Check(test.input) != test.result {
			t.Errorf("Test data:[%v],[%v] not in source slice.", index, test.input)
		}
	}
}
