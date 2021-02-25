package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	res := Sum(2, 3)

	if res != 5 {
		t.Error("mySum result wrong expected: 5 got", res)
	}
}

func TestSumTable(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{1, 2}, 3},
		{[]int{2, 4, 6}, 12},
		{[]int{-1, 0, 1}, 0},
		{[]int{21, 21}, 42},
		{[]int{0}, 0},
	}

	for _, v := range tests {
		res := Sum(v.data...)

		if res != v.answer {
			t.Error("mySum result wrong expected:", v.answer, " got", res)
		}
	}
}

func ExampleSum() {
	fmt.Println(Sum(5, 5))
	//Output:
	//10
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(2, 2)
	}
}
