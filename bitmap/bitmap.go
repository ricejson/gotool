package bitmap

import (
	"errors"
	"fmt"
)

const (
	DefaultCapacity = 128
	ByteSize        = 8
)

// BitMap 位图
type BitMap struct {
	data     []byte // 底层数组
	usedSize int    // 存储的元素个数
	capacity int    // 容量（单位：字节）

}

func NewBitMap() *BitMap {
	data := make([]byte, DefaultCapacity)
	return &BitMap{
		data:     data,
		usedSize: 0,
		capacity: DefaultCapacity,
	}
}

// ensureCapacity 自动扩容
func (bitMap *BitMap) ensureCapacity(elem int) error {
	// 计算所需的字节数
	var needCapacity = (elem / ByteSize) + 1
	if needCapacity > bitMap.capacity {
		// 进行扩容
		var newData = make([]byte, needCapacity)
		copy(newData, bitMap.data)
		bitMap.data = newData
		bitMap.capacity = needCapacity
	}
	return nil
}

// Set 添加元素
func (bitMap *BitMap) Set(elem int) error {
	// 校验元素合法性
	if elem < 0 {
		return errors.New("element must be greater than zero")
	}
	err := bitMap.ensureCapacity(elem)
	if err != nil {
		return fmt.Errorf("set element fail：%w", err)
	}
	// 计算bit位置
	var byteIdx = elem / ByteSize
	var bitIdx = elem % ByteSize
	// 添加元素
	bitMap.data[byteIdx] |= 1 << bitIdx
	bitMap.usedSize++
	return nil
}

// Contains 判断元素是否存在
func (bitMap *BitMap) Contains(elem int) (bool, error) {
	// 校验元素合法性
	if elem < 0 {
		return false, errors.New("element must be greater than zero")
	}
	// 校验元素是否超出容量
	var byteIdx = elem / ByteSize
	if byteIdx >= bitMap.capacity {
		return false, nil
	}
	var bitIdx = elem % ByteSize
	return (bitMap.data[byteIdx]>>bitIdx)&1 == 1, nil
}

// Reset 重置元素
func (bitMap *BitMap) Reset(elem int) error {
	// 查询该元素是否已经存在
	exists, err := bitMap.Contains(elem)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	// 重置值
	var byteIdx = elem / ByteSize
	var bitIdx = elem % ByteSize
	bitMap.data[byteIdx] ^= 1 << bitIdx
	bitMap.usedSize--
	return nil
}

// Size 获取已经存储的元素个数
func (bitMap *BitMap) Size() int {
	return bitMap.usedSize
}
