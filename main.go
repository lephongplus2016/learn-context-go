/*
	context can be used to pass value between layter
*/

package main

import (
	"context"
	"fmt"
	"time"
)

// store data key-value in context
func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "123")
}

// get value id in context
func getVarID(ctx context.Context) {
	varID := ctx.Value("request-id")
	fmt.Println("varID:", varID)

	for {
		select {
		case <-ctx.Done(): // after 2s-> out
			fmt.Println("time out")
			return
		default:
			fmt.Println("working")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// ctx := context.Background()

	// - tạo ra 1 context mới
	// - có deadline là 2s, nếu vượt qua thì sẽ close ctx
	// -> ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // hủy hoãn cancel thì ctx mới tồn tại

	ctx = enrichContext(ctx)

	// chạy ngầm go rountine
	go getVarID(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("oh no, I've exceeded the deadline")
	}

	time.Sleep(2 * time.Second)
}
