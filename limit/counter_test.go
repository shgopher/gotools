package limit

import (
	"fmt"
	"testing"
	"time"
)

func TestConter(t *testing.T) {
	c := NewConter(2, time.Second)
	for i:= 0;i < 10;i++ {
		go func() {
			for {
				if c.Run(){
					time.Sleep(time.Second*1)
					fmt.Println("runing")
				}else {
					time.Sleep(time.Second*2)
					fmt.Println("waiting")
				}
			}
		}()
	}
time.Sleep(time.Second*100)
}
