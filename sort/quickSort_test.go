package sort

import (
	"practice/test"
	"testing"
)

func TestQuickSort(t *testing.T) {
	test.RunSort(t, QuickSort, test.DefaultParams)
}