package concurency

import (
	"fmt"
	"sync"
)

func SumOfSlice(sl []int, wg *sync.WaitGroup, id int, ch chan int) {
	sumOfElements := 0
	for _, value := range sl {
		sumOfElements += value
	}
	ch <- sumOfElements
	fmt.Printf("Сумма элементов %d горутины: %d\n", id, sumOfElements)
	wg.Done()
}
