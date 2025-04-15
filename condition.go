package gotool

// IfThenElse 根据条件返回对应的泛型结果
// 如果condition成立返回trueValue，反之返回falseValue
func IfThenElse[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
