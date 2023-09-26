package pipline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func ArraySource(a ...int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		// read to memory
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		fmt.Printf("read done: %v\n", time.Now().Sub(startTime))

		// 排序
		sort.Ints(a)
		fmt.Printf("InMemSort done: %v\n", time.Now().Sub(startTime))

		// 输出
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Printf("Merge done: %v\n", time.Now().Sub(startTime))
	}()
	return out
}

// chunkSize 分块大小
func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		byteRead := 0
		for {
			n, err := reader.Read(buffer)
			byteRead += n
			if n > 0 {
				v := int(
					binary.BigEndian.Uint64(buffer))
				out <- v
			}

			if err != nil || (chunkSize != -1 && byteRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriteSink(w io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		w.Write(buffer)
	}
}

// 生成随机数
func RandomSource(num int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		for i := 0; i < num; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}

	m := len(inputs) / 2
	// merge inputs[0..m] and inputs [m..end]
	return Merge(
		MergeN(inputs[:m]...),
		MergeN(inputs[m:]...))
}
