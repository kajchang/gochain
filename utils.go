package gochain

import (
	"encoding/binary"
	"math"
	"strings"
)

func EncodeFloat64(f float64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, math.Float64bits(f))
	return buf
}

func EncodeUint64(i uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

func MiddlePadString(s string, l int) string {
	padLength := l - len(s)
	rightPadLength, leftPadLength := padLength / 2, padLength / 2
	if padLength % 2 != 0 {
		leftPadLength++
	}
	return strings.Repeat(" ", leftPadLength) + s + strings.Repeat(" ", rightPadLength)
}
