package sort

import (
	"practice/test"
	"testing"
)

func TestHeapSort(t *testing.T) {
	test.RunSort(t, HeapSort, test.DefaultParams)
}
