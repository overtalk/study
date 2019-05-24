package main

import (
	"context"
	"fmt"
)

var (
	key1 = "key1"
	key2 = "key2"
	key3 = "key3"
)

func testMuli() {
	ctx1 := context.WithValue(context.Background(), key1, "1")
	ctx2 := context.WithValue(ctx1, key2, "2")
	ctx3 := context.WithValue(ctx2, key3, "3")

	value1 := ctx3.Value(key1)
	value2 := ctx3.Value(key2)
	value3 := ctx3.Value(key3)

	fmt.Println(value1, value2, value3)
}
