package id

import (
	"fmt"
	"sync"
	"testing"
)

var (
	count = 100000
)

func TestNewSnowFlake(t *testing.T) {
	snowFlake(10)
}
func TestNewSnowFlake2(t *testing.T) {
	wa := new(sync.WaitGroup)
	wa.Add(10)
	for i := 0;i< 10;i++ {
		t := int64(i)
		go func(t int64) {
			defer wa.Done()
			snowFlake(t)
		}(t)
	}
	wa.Wait()

}

func snowFlake(wrokid int64)  {
	sf, err := NewSnowFlake(wrokid)
	if err != nil {
		fmt.Println(err)
	}
	ch := make(chan int64)
	ma := make(map[int64]int)
	for i := 0; i < count; i++ {
		go func() {
			ch <- sf.GetID()
		}()
	}
	L:
	for i := 0; i < count; i++ {
		select {
		case d := <-ch:
			if _,ok := ma[d];ok {
				fmt.Println("bad")
				break L
			}
			ma[d]++
		default:
			break L
		}
	}
}