package slice

// Find 在src中查找符合条件的第一个元素
// 如果找到了返回值本身和true
// 如果没找到返回类型空值以及false
func Find[T any](src []T, matchFunc func(T) bool) (T, bool) {
	for _, v := range src {
		if matchFunc(v) {
			return v, true
		}
	}
	var zeroValue T
	return zeroValue, false
}

// FindAll 在src中查找符合条件的所有元素并返回切片
func FindAll[T any](src []T, matchFunc func(T) bool) []T {
	var ret = make([]T, 0, len(src))
	for _, v := range src {
		if matchFunc(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
