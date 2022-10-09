package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func DifficultFuncCtxStart() {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		DifficultFuncCtx(ctx)
	}()
	wg.Wait()
}

func DifficultFuncCtx(ctx context.Context) {
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err().Error())
			return
		default:
		}

		difficultStepCtx()
	}
	fmt.Println("DifficultFunc all step finish")
}

func difficultStepCtx() {
	time.Sleep(2 * time.Second)
	fmt.Println("difficult step finish")
}
