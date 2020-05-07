package limit

import (
	"fmt"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	tb := NewTokenBucket(10,time.Minute)
	for i := 0;i < 2;i++ {
		go func() {
			for {
				tb.Run()
				fmt.Println("runing")
			}
		}()
	}
	time.Sleep(time.Second*100)
}