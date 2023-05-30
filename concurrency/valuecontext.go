package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9) //컨텍스트에 값을 추가한다
	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil { //컨텍스트에서 값을 읽어온다.
		n := v.(int)
		fmt.Printf("Square:%d", n*n)
	}
	wg.Done()
}
