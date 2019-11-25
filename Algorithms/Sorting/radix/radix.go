package radix

func maxBit(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	bit := 1
	for max >= 10 {
		// divide and assign
		max /= 10
		bit++
	}
	return bit
}

func RadixSort(arr []int) {
	mBit := maxBit(arr)
	tmp := make([]int, len(arr), len(arr))
	count := make([]int, 10, 10)
	for i, radix := 1, 1; i < mBit; i++ {
		for j := 0; j < 10; j++ {
			count[j] = 0
		}
		for j := 0; j < len(arr); j++ {
			k := (arr[j] / radix) % 10
			count[k]++
		}
		for j := 1; j < 10; j++ {
			count[j] = count[j-1] + count[j]
		}
		for j := len(arr) - 1; j >= 0; j-- {
			k := (arr[j] / radix) % 10
			tmp[count[k]-1] = arr[j]
			count[k]--
		}
		copy(arr, tmp)
		radix *= 10
	}
}
