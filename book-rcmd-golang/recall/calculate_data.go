package recall

import (
	"context"
	"fmt"
	"sync/atomic"
)

var count int32

func Init(ctx context.Context) {

	for i := 0; i < 12; i++ {
		go CalculateExecutor(ctx, InputDataChannel)
	}

}

func CalculateExecutor(ctx context.Context, channel chan InputData) {
	for {
		select {
		case <-ctx.Done():
			break
		case data := <-channel:
			calculate(data)
		}
	}
}

func calculate(data InputData) {
	atomic.AddInt32(&count, 1)
	if count%10000 == 0 {
		fmt.Println(count)
	}
}
