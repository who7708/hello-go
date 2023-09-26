package main

import (
	"encoding/hex"
	"fmt"

	"github.com/spaolacci/murmur3"
)

// 打印 murmur3 hash 后的字符串，与java一致
func murmur32ToString(str string) string {
	b1 := []byte(str)
	hasher := murmur3.New32()
	hasher.Write(b1)
	r := hasher.Sum([]byte{})
	// for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
	// 	r[i], r[j] = r[j], r[i]
	// }
	r = reverse(r)

	// fmt.Printf("r: %v\n", string(r))
	s := hex.EncodeToString(r)
	fmt.Printf("murmur3(%s): %v\n", str, s)
	return str
}

// []byte 反转
func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
