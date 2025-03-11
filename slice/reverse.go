package slice

// Reverse 翻转src切片中的元素并返回一个新切片
func Reverse[T any](src []T) []T {
	var n = len(src)
	var ret = make([]T, n)
	for i := 0; i < n; i++ {
		ret[i] = src[n-i-1]
	}
	return ret
}

// ReverseSelf 原地翻转src切片中的元素
func ReverseSelf[T any](src []T) {
	left, right := 0, len(src)-1
	for left < right {
		src[left], src[right] = src[right], src[left]
		left++
		right--
	}
}
