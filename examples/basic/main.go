package main

import (
	"fmt"
	lib "github.com/1Pas2/pas2-go-token-bucket"
	"time"
)

func main() {
	now := time.Now()
	bucket, err := lib.New(10, 2, now)
	if err != nil {
		panic(err)
	}
	fmt.Println(bucket.Take(now, 1))
}
