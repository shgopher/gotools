package limit

import (
	"fmt"
	"testing"
	"time"
)

func TestConter(t *testing.T) {
	c := NewConter(20, time.Second)
	for i:= 0;i < 10;i++ {
		go func() {
			for {
				if c.Run(){
					time.Sleep(time.Second/10)
					fmt.Println("runing")
				}else {
					time.Sleep(time.Second/10)
					fmt.Println("waiting")
				}
			}
		}()
	}
time.Sleep(time.Second*50)
}
