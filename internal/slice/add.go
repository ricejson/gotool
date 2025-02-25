package slice

import (
	"errors"
)

// Add 往切片指定索引处添加元素
// 允许范围是[0, len(s))
func Add[T any](src []T, index int, element T) (newSlice []T, err error) {
	length := len(src)
	if index < 0 || index > length {
		return nil, errors.New("index out of range")
	}
	var zeroValue T
	newSlice = append(src, zeroValue)
	for i := length - 1; i >= index; i-- {
		newSlice[i+1] = newSlice[i]
	}
	newSlice[index] = element
	return newSlice, nil
}
