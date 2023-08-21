package pipline

import "sort"

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
		a := []int{}
		for v := range in {
			a = append(a, v)
		}

		sort.Ints(a)

		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}
