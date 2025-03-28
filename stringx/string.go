package stringx

import "unicode"

// IsEmpty 判断字符串是否为空
// 如果字符串为 "" 则返回true，否则返回false
func IsEmpty(str string) bool {
	if str == "" {
		return true
	}
	return false
}

// IsNotEmpty 判断字符串是否不为空
// 如果字符串为 "" 则返回false，否则返回true
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// IsBlank 判断字符串是否全为空白字符
// 如果字符串全为空白字符 则返回true，否则返回false
func IsBlank(str string) bool {
	for _, v := range str {
		if !unicode.IsSpace(v) {
			return false
		}
	}
	return true
}

// IsNotBlank 判断字符串是否不全为空白字符
// 如果字符串不全为空白字符 则返回true，否则返回false
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}
