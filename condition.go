package gotool

// IfThenElse 根据条件返回对应的泛型结果
// 如果condition成立返回trueValue，反之返回falseValue
func IfThenElse[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

// IfThenElseFunc 根据条件执行对应的函数逻辑并返回泛型结果
// 如果condition成立执行trueFunc并返回，反之执行falseFunc并返回
func IfThenElseFunc[T any](condition bool, trueFunc, falseFunc func() (T, error)) (T, error) {
	if condition {
		return trueFunc()
	}
	return falseFunc()
}
