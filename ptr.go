package gotool

// ToPtr 获取对应类型的指针类型
func ToPtr[T any](src T) *T {
	return &src
}
