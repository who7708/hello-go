package pipline

import (
	"encoding/binary"
	"io"
	"math/rand"
	"sort"
)

func ArraySource(a ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// read to memory
		a := []int{}
		for v := range in {
			a = append(a, v)
		}

		// 排序
		sort.Ints(a)

		// 输出
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
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
	}()
	return out
}

func ReaderSource(reader io.Reader) <-chan int {
	out := make(chan int)
	go func() {
		buffer := make([]byte, 8)
		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				v := int(
					binary.BigEndian.Uint64(buffer))
				out <- v
			}

			if err != nil {
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
	out := make(chan int)

	go func() {
		for i := 0; i < num; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}
