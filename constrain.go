package gotool

// RealNumber 实数类型约束
type RealNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Number 数值类型约束
type Number interface {
	RealNumber | ~complex64 | ~complex128
}
