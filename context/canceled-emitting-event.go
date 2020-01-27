// Emitting a cancellation event
package context

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func operation1() error {
	//  sleep
	time.Sleep(10 * time.Second)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	select {
	case <-time.After(50 * time.Second):
		fmt.Println("finished")
	case <-ctx.Done():
		fmt.Println("interupt")
	}
}

func main() {
	ctx := context.Background()
	// 如果operation1失败了，则operation2退出执行
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		err := operation1()
		if err != nil {
			cancel()
		}
	}()
	
	// i := 4 
	// i += 1 

	fmt.Println("go context")
	operation2(ctx)
}
