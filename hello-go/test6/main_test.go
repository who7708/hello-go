package main

import (
	"testing"
)

func BenchmarkMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5Hash()
	}
}

func BenchmarkSha1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha1Hash()
	}
}

func BenchmarkCrc32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		crc32Hash()
	}
}

func BenchmarkMur32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		murmur32Hash()
	}
}

func BenchmarkMur64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		murmur64Hash()
	}
}
