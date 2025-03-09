package slice

// IndexOf 返回在src当中第一个与dst相同的下标(接收comparable类型)
// 如果没找到返回-1
func IndexOf[T comparable](src []T, dst T) int {
	return IndexOfFunc[T](src, dst, func(e1, e2 T) bool {
		return e1 == e2
	})
}

// IndexOfFunc 返回在src当中第一个与dst相同的下标(接收任意类型)
// 如果没找到返回-1
func IndexOfFunc[T any](src []T, dst T, compareTo func(e1, e2 T) bool) int {
	for i, v := range src {
		if compareTo(v, dst) {
			return i
		}
	}
	return -1
}

// LastIndexOf 返回在src当中最后一个与dst相同的下标(接收comparable类型)
// 如果没找到返回-1
func LastIndexOf[T comparable](src []T, dst T) int {
	return LastIndexOfFunc[T](src, dst, func(e1, e2 T) bool {
		return e1 == e2
	})
}

// LastIndexOfFunc 返回在src当中最后一个与dst相同的下标(接收任意类型)
// 如果没找到返回-1
func LastIndexOfFunc[T any](src []T, dst T, compareTo func(e1, e2 T) bool) int {
	for i := len(src) - 1; i >= 0; i-- {
		if compareTo(src[i], dst) {
			return i
		}
	}
	return -1
}

// IndexOfAll 返回在src当中所有与dst相同的下标(接收comparable类型)
// 如果没找到返回空切片
func IndexOfAll[T comparable](src []T, dst T) []int {
	return IndexOfAllFunc[T](src, dst, func(e1, e2 T) bool {
		return e1 == e2
	})
}

// IndexOfAllFunc 返回在src当中所有与dst相同的下标(接收任意类型)
// 如果没找到返回空切片
func IndexOfAllFunc[T any](src []T, dst T, compareTo func(e1, e2 T) bool) []int {
	var indexes = make([]int, 0, len(src))
	for i, v := range src {
		if compareTo(v, dst) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
