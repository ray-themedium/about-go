package main

import (
	"context"
	"fmt"
)

type aKey string

func searchKey(ctx context.Context, k aKey) {
	v := ctx.Value(k)
	if v != nil {
		fmt.Println("found value:", v)
		return
	} else {
		fmt.Println("key not found:", k)
	}
}

func main() {
	myKey := aKey("testKey")
	ctx := context.WithValue(context.Background(), myKey, "testValue")

	searchKey(ctx, myKey)
	searchKey(ctx, aKey("badKey"))
	emptyCtx := context.TODO()
	searchKey(emptyCtx, aKey("badKey"))
}
