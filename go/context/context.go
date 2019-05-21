package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var key interface{} = "index"
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) {
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("go done ,value = ", ctx.Value(key))
					return // returning not to leak the goroutine
				default:
					time.Sleep(time.Second)
				}
			}
		}()

	}

	ctx, cancel := context.WithCancel(context.Background())

	// var subCtxArr []context.Context
	for i := 0; i < 10; i++ {
		// subCtxArr = append(subCtxArr, context.WithValue(ctx, key, i))
		go gen(context.WithValue(ctx, key, i))
	}

	time.Sleep(time.Second)

	cancel() // cancel when we are finished consuming integers

	time.Sleep(3 * time.Second)
	fmt.Println("main is over")
}
