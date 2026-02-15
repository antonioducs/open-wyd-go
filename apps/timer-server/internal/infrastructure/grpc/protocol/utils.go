package protocol

import (
	"unsafe"
)

func Cast[T any](data []byte) *T {
	return (*T)(unsafe.Pointer(&data[0]))
}

func CopyTextToByte(length int, src string) []byte {
	dst := make([]byte, length)
	copy(dst, src)
	return dst
}
