package test

import (
	"math/rand"
	"testing"
	"time"
)

type Params struct {
	Cases      int
	RangeStart int
	RangeEnd   int
	Length     int
}

var DefaultParams Params = Params{10000, -500, 500, 900}

func verifySorted(t *testing.T, arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			t.Fail()
		}
	}
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func GenerateArr(params Params) []int {
	arr := make([]int, params.Length)
	for i := range arr {
		arr[i] = rand.Intn(params.RangeEnd-params.RangeStart) + params.RangeStart
	}
	return arr
}

func RunSort(t *testing.T, sort func([]int), params Params) {
	fn := func() {
		var arr []int

		setups := []func(){
			// Randomized (generated in random order)
			func() {
				arr = GenerateArr(params)
			},
			// Reversed
			func() {
				reverse(arr)
			},
			// In-order (already sorted from last test)
			func() {},
			// Single element
			func() {
				i := rand.Intn(len(arr) - 1)
				arr = arr[i : i+1]
			},
			// Empty
			func() {
				arr = arr[:0]
			},
		}

		for _, setup := range setups {
			setup()
			sort(arr)
			verifySorted(t, arr)
		}

	}
	Run(fn, params)
}

func Run(f func(), params Params) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < params.Cases; i++ {
		f()
	}
}
