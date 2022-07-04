package sort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr) - 1)
}

func quickSort(arr []int, left, right int) {
	if (left < right) {
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot - 1)
		quickSort(arr, pivot + 1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	swap := left

	for i := left; i < right; i++ {
		if arr[i] < pivot {
			arr[i], arr[swap] = arr[swap], arr[i]
			swap++
		}
	}

	arr[swap], arr[right] = pivot, arr[swap]
	return swap
}