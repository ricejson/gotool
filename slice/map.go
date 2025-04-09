package slice

// Map 将[]Src切片转换为[]Dst
func Map[Src any, Dst any](s []Src, f func(Src) Dst) []Dst {
	var ret = make([]Dst, len(s))
	for i, v := range s {
		ret[i] = f(v)
	}
	return ret
}

// Filter 将原切片不符合要求的元素过滤并转换为新切片返回
func Filter[T any](src []T, matchFunc func(e T) bool) []T {
	var ret = make([]T, 0, len(src))
	for _, v := range src {
		if matchFunc(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
