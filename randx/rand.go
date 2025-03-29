package randx

import "errors"

var (
	a uint32 = 1664525    // 乘数
	c uint32 = 1013904223 // 增量
	m        = 4294967296 // 模数（2^32）
)

// LCG的seed
var seed uint32

func uint32LCG() uint32 {
	seed = uint32(int(a*seed+c) % m)
	return seed
}

func intLCG() int {
	return int(uint32LCG())
}

// Int 随机生成[minInclude, maxExclude)的int类型
func Int(minInclude, maxExclude int) (int, error) {
	// 处理最小数大于最大数的情况
	if minInclude >= maxExclude {
		return 0, errors.New("minInclude >= maxExclude")
	}
	// 处理负数的情况
	if minInclude < 0 {
		return 0, errors.New("minInclude < 0")
	}
	rangeSize := maxExclude - minInclude
	return minInclude + intLCG()%rangeSize, nil
}
