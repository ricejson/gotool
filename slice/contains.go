package slice

// Contains 判断src当中是否存在dst元素(接收comparable类型)
func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc[T](src, dst, func(e1, e2 T) bool {
		return e1 == e2
	})
}

// ContainsFunc 判断src当中是否存在dst元素(接收任意类型)
func ContainsFunc[T any](src []T, dst T, compareTo func(e1, e2 T) bool) bool {
	for _, v := range src {
		if compareTo(v, dst) {
			return true
		}
	}
	return false
}

// ContainsAny 判断src当中是否存在dst中的任一元素(接收comparable类型)
func ContainsAny[T comparable](src []T, dst []T) bool {
	return ContainsAnyFunc[T](src, dst, func(e1, e2 T) bool {
		return e1 == e2
	})
}

// ContainsAnyFunc 判断src当中是否存在dst中的任一元素(接收任意类型)
func ContainsAnyFunc[T any](src []T, dst []T, compareTo func(e1, e2 T) bool) bool {
	for _, v := range dst {
		if ContainsFunc[T](src, v, compareTo) {
			return true
		}
	}
	return false
}

// ContainsAll 判断src当中是否存在dst中的所有元素(接收comparable类型)
func ContainsAll[T comparable](src []T, dst []T) bool {
	return ContainsAllFunc[T](src, dst, func(e1, e2 T) bool {
		return e1 == e2
	})
}

// ContainsAllFunc 判断src当中是否存在dst中的所有元素(接收任意类型)
func ContainsAllFunc[T any](src []T, dst []T, compareTo func(e1, e2 T) bool) bool {
	for _, v := range dst {
		if !ContainsFunc[T](src, v, compareTo) {
			return false
		}
	}
	return true
}
