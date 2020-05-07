package limit

import (
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	tb := NewTokenBucket(1,time.Second)
	for i := 0;i < 2;i++ {
		go func() {
			for {
				tb.Run()
				t.Log("runing")
			}
		}()
	}
	time.Sleep(time.Second*50)
}