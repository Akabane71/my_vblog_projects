package models

import (
	"fmt"
	"testing"
	"unsafe"
)

type s1 struct {
	a int8
	b string
	c int8
}

type s2 struct {
	a int8
	b int8
	c string
}

// 内存对齐 示例
func TestStruct(t *testing.T) {
	v1 := s1{
		a: 1,
		b: "hello",
		c: 1,
	} // 内存占用 32

	v2 := s2{
		a: 1,
		b: 2,
		c: "hello",
	} // 内存占用 24
	fmt.Println(unsafe.Sizeof(v1), unsafe.Sizeof(v2))
}
