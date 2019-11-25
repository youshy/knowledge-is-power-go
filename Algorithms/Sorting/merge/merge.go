package merge

func merge(left, right []int) []int {
	lenL, lenR := len(left), len(right)
	res := make([]int, 0, lenL+lenR)
	i, j := 0, 0

	for i < lenL && j < lenR {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	// tail
	if i < lenL {
		res = append(res, left[i:]...)
	}

	if j < lenR {
		res = append(res, right[j:]...)
	}

	return res
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2 // array split point
	return merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}
