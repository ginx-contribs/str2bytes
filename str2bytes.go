package str2bytes

import "unsafe"

// Str2Bytes return []byte representation of string, make sure underlying data will not be modified.
func Str2Bytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// Bytes2Str return string representation of []byte
func Bytes2Str(bs []byte) string {
	return unsafe.String(unsafe.SliceData(bs), len(bs))
}
