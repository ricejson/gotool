package slice

// Map 将[]Src切片转换为[]Dst
func Map[Src any, Dst any](s []Src, f func(Src) Dst) []Dst {
	var ret = make([]Dst, len(s))
	for i, v := range s {
		ret[i] = f(v)
	}
	return ret
}
