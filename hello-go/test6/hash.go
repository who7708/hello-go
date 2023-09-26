package main

import (
	"crypto/md5"
	"crypto/sha1"
	"hash/crc32"

	"github.com/spaolacci/murmur3"
)

var str = "hello world"

func md5Hash() [16]byte {
	return md5.Sum([]byte(str))
}

func sha1Hash() [20]byte {
	return sha1.Sum([]byte(str))
}

func crc32Hash() uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

func murmur32Hash() uint32 {
	return murmur3.Sum32([]byte(str))
}

func murmur64Hash() uint64 {
	return murmur3.Sum64([]byte(str))
}
