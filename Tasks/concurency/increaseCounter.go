package concurency

import (
	"fmt"
	"sync"
)

func increaseNumber(counter *int, wg *sync.WaitGroup, mu *sync.Mutex, id int) {
	defer wg.Done()
	mu.Lock()
	*counter += 1
	fmt.Printf("Горутина %d завершила работу. Число равно: %d\n", id, *counter)
	mu.Unlock()
}
