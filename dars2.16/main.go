package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("go context")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go myfunc(ctx)
	select{
	case <-ctx.Done():
		fmt.Println("oh no")
	}
	time.Sleep(2*time.Second)
	context.WithoutCancel()
}

func myfunc(ctx context.Context)  {
	for{
		select{
		case <-ctx.Done():
			fmt.Println("vaht tugadi")
			return
		default:
			fmt.Println("hello")
		}
		time.Sleep(500*time.Millisecond)
	}
}

