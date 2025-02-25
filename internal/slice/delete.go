package slice

import "errors"

// Delete 从切片指定索引处删除元素
// 允许范围是[0, len(s))
func Delete[T any](src []T, index int) (newSlice []T, element T, err error) {
	length := len(src)
	if index < 0 || index >= length {
		var zeroValue T
		return nil, zeroValue, errors.New("index out of range")
	}
	// TODO: 待实现缩容功能
	var target = src[index]
	for i := index; i < length-1; i++ {
		src[index] = src[index+1]
	}
	src = src[:length-1]
	return src, target, nil
}
