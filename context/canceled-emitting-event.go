// Emitting a cancellation event
package main

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

// explain code below here
// 1. context.Background() 会返回一个空的Context，这个空的Context一般用于整个Context树的根节点。
// 2. context.WithCancel(parent) 会返回一个子类Context，这个子类Context的父类为parent，当父类的Context被取消时，该Context也会被取消。
// 3. context.WithDeadline(parent, time) 会返回一个子类Context，这个子类Context的父类为parent，当到达指定时间时，该Context会被取消。
// 4. context.WithTimeout(parent, time) 和WithDeadline类似，二者区别是，WithTimeout是计算相对时间，到达指定时间后，该Context会被取消。
// 5. context.WithValue(parent, key, value) 会返回一个子类Context，这个子类Context的父类为parent，同时在该Context中绑定了一个键值对，可以通过Value(key)来获取对应的值，当然，该值只能是线程安全的。


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

	fmt.Println("go context")
	operation2(ctx)
}
