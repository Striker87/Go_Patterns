package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
Шаблон Circuit Breaker (Размыкатель цепи) шаблон стабильности автоматически отключает сервисные функции в ответ на вероятную неисправность, чтобы предотвратить более крупные или каскадные отказы, устранить повторяющиеся ошибки и обеспечить разумную реакцию на ошибки.
*/

type Circuit func(ctx context.Context) (*sqlx.DB, error)

func main() {
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		fn := Breaker(conn, 10)
		res, err := fn(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("res:", res)
	}
}

func conn(ctx context.Context) (*sqlx.DB, error) {
	fmt.Println("conn")
	return sqlx.Connect("mysql", "")
}

func Breaker(conn Circuit, failureThreshold uint) Circuit {
	var consecutiveFailures = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex
	fmt.Println(1)

	return func(ctx context.Context) (*sqlx.DB, error) {
		fmt.Println(2)
		m.RLock() // Установить "блокировку чтения"
		d := consecutiveFailures - int(failureThreshold)
		fmt.Println("d:", d)

		if d >= 0 {
			fmt.Println(3)
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return nil, errors.New("service unreachable")
			}
		}
		fmt.Println(11)
		m.RUnlock()                // Освободить блокировку чтения
		response, err := conn(ctx) // Послать запрос, как обычно
		m.Lock()                   // Заблокировать общие ресурсы
		defer m.Unlock()
		lastAttempt = time.Now() // Зафиксировать время попытки

		if err != nil { // Если Circuit вернула ошибку,
			fmt.Println(5)

			consecutiveFailures++ // увеличить счетчик ошибок
			fmt.Println("consecutiveFailures:", consecutiveFailures)
			return response, err // и вернуть ошибку
		}

		fmt.Println(6)
		consecutiveFailures = 0 // Сбросить счетчик ошибок
		return response, nil
	}
}
