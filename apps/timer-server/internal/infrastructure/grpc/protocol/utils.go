package protocol

import "unsafe"

func Cast[T any](data []byte) *T {
	return (*T)(unsafe.Pointer(&data[0]))
}
