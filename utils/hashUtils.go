package utils

import (
	"bytes"
	"encoding/gob"

	"hash/crc32"
)

// String hashes a string to a unique hashcode.

// Hashcode crc32 returns a uint32, but for our use we need
// and non negative integer. Here we cast to an integer
// and invert it if the result is negative.
func Hashcode(b []byte) int {
	v := int(crc32.ChecksumIEEE(b))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func CheckKeyEquality(a, b []byte) bool {
	isEqual := true

	if len(a) != len(b) {
		return false
	}

	for index, item := range a {
		if item != b[index] {
			isEqual = false
			return isEqual
		}
	}
	return isEqual
}

func Encode(obj interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(obj); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Hash(h, len int) int {
	return h & (len - 1)
}
