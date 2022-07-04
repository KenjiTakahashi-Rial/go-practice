package algorithms

import (
	"practice/test"
	"testing"
)

func verifyUnsorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return true
		}
	}
	return false
}

func TestShuffle(t *testing.T) {
	params := test.DefaultParams
	f := func() {
		arr := test.GenerateArr(params)
		Shuffle(arr)
		if !verifyUnsorted(arr) {
			t.Fail()
		}
	}
	test.Run(f, params)
}
