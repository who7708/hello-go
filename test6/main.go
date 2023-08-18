package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/spaolacci/murmur3"
)

type Student struct {
	Name     string `json:"name"`
	Age      uint8  `json:"age"`
	Location string `json:"location"`
}

func main() {

	// testMurmur3("hello1")
	// testMurmur3("hello2")
	// testMurmur3("hello3")

	// fmt.Println("======================")

	testMurmur3_2("hello1")
	testMurmur3_2("hello2")
	testMurmur3_2("hello3")

	// fmt.Println("======================")
	// testHash()
}

// 打印 murmur3 hash 后的字符串，与java一致
func testMurmur3_2(str string) string {
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

// 打印 murmur3 hash 后的字符串，与java一致
func testMurmur3(str string) string {
	b1 := []byte(str)

	_hash := murmur3.Sum32(b1)

	bs := []byte{byte(_hash), byte(_hash >> 8), byte(_hash >> 16), byte(_hash >> 24)}

	s := hex.EncodeToString(bs)

	fmt.Printf("string(b): %v\n", s)
	return s
}

func testHash() {

	stu := Student{
		Name:     "zhangsan",
		Age:      23,
		Location: "shanghai",
	}

	// r, _ := sonic.MarshalString(&stu)
	r, _ := sonic.Marshal(&stu)
	var sampleInt int32 = 1

	sampleString := fmt.Sprint(sampleInt)
	fmt.Printf("%+v %+v\n", sampleInt, sampleString)
	fmt.Println("=====================================")

	hash1 := murmur3.Sum32(r)
	fmt.Printf("string(hash): %v\n", hash1)
	fmt.Printf("string(hash): %v\n", fmt.Sprint(hash1))
	fmt.Println("=====================================")

	b5 := murmur3.New32().Sum(r)
	fmt.Printf("string(b5): %v\n", string(b5))
	fmt.Printf("murmur3.New32().Sum(r): %v\n", base64.RawStdEncoding.EncodeToString(b5[:]))
	fmt.Printf("murmur3.New32().Sum(r): %v\n", hex.EncodeToString(b5[:]))
	fmt.Println("=====================================")

	s := strconv.FormatUint(uint64(hash1), 10)
	fmt.Printf("s: %v\n", s)
	fmt.Println("=====================================")

	fmt.Printf("murmur3.New32().Sum(\"hello1\"): %v\n", string(murmur3.New32().Sum([]byte("hello1"))))
	fmt.Printf("murmur3.New32().Sum(\"hello2\"): %v\n", murmur3.New32().Sum([]byte("hello2")))
	fmt.Printf("murmur3.New32().Sum(\"hello3\"): %v\n", murmur3.New32().Sum([]byte("hello3")))
	fmt.Println("=====================================")

	b := md5.Sum(r)
	s2 := base64.RawStdEncoding.EncodeToString(b[:])
	s3 := hex.EncodeToString(b[:])
	fmt.Printf("md5.Sum(r): %v\n", s2)
	fmt.Printf("md5.Sum(r): %v\n", s3)
	fmt.Println("=====================================")

	b2 := sha1.Sum(r)
	fmt.Printf("sha1.Sum(r): %v\n", base64.RawStdEncoding.EncodeToString(b2[:]))
	fmt.Printf("sha1.Sum(r): %v\n", hex.EncodeToString(b2[:]))
	fmt.Println("=====================================")

	b3 := sha256.Sum256(r)
	fmt.Printf("sha256.Sum(r): %v\n", base64.RawStdEncoding.EncodeToString(b3[:]))
	fmt.Printf("sha256.Sum(r): %v\n", hex.EncodeToString(b3[:]))
	fmt.Println("=====================================")

	b4 := sha512.Sum512(r)
	fmt.Printf("sha512.Sum(r): %v\n", base64.RawStdEncoding.EncodeToString(b4[:]))
	fmt.Printf("sha512.Sum(r): %v\n", hex.EncodeToString(b4[:]))
	fmt.Println("=====================================")
}

// func murmur32Hash() uint32 {
// 	return murmur3.Sum32([]byte(str))
// }
