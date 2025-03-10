package slice

import (
	"errors"
	"github.com/ricejson/gotool"
)

// Max 返回切片当中的最大值
// 如果方法传递空切片或者nil则会返回T类型零值和特定错误
func Max[T gotool.RealNumber](slice []T) (T, error) {
	var maxVal T
	if slice == nil {
		return maxVal, errors.New("slice is nil")
	}
	var n = len(slice)
	if n == 0 {
		return maxVal, errors.New("slice is empty")
	}
	maxVal = slice[0]
	for i := 1; i < n; i++ {
		if slice[i] > maxVal {
			maxVal = slice[i]
		}
	}
	return maxVal, nil
}

// Min 返回切片当中的最小值
// 如果方法传递空切片或者nil则会返回T类型零值和特定错误
func Min[T gotool.RealNumber](slice []T) (T, error) {
	var minVal T
	if slice == nil {
		return minVal, errors.New("slice is nil")
	}
	var n = len(slice)
	if n == 0 {
		return minVal, errors.New("slice is empty")
	}
	minVal = slice[0]
	for i := 1; i < n; i++ {
		if slice[i] < minVal {
			minVal = slice[i]
		}
	}
	return minVal, nil
}

// Sum 返回切片求和结果
// 如果方法传递空切片或者nil则会返回T类型零值和特定错误
func Sum[T gotool.Number](slice []T) (T, error) {
	var sumVal T
	if slice == nil {
		return sumVal, errors.New("slice is nil")
	}
	var n = len(slice)
	for i := 0; i < n; i++ {
		sumVal += slice[i]
	}
	return sumVal, nil
}
