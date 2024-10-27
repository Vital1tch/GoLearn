package concurency

import (
	"context"
	"fmt"
	"time"
)

func processTask(ctx context.Context) {
	for i := 0; i < 6; i++ {
		select {
		case <-ctx.Done(): //Контекст завершает дедлайн по истечению 4 секунд
			fmt.Println("Дедлайн провален")
			return
		default: //Будет работать пока дедлайн не кончится каждую секунду
			time.Sleep(1 * time.Second)
			fmt.Println("Выполнение задачи....")
		}
	}
	fmt.Println("Задача завершена до дедлайна!") //Если цикл отработал 6 раз в секунду до дедлайна
}
