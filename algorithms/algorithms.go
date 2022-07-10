package algorithms

import (
	"math/rand"
	"time"
)

func Shuffle[T any](arr []T) {
	rand.Seed(time.Now().UnixNano())
	n := len(arr)
	for i := 0; i < n-1; i++ {
		j := i + rand.Intn(n-i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
