package id

import (
	"fmt"
	"sync"
	"testing"
)

var (
	count = 10000
)

func TestNewSnowFlake(t *testing.T) {
	snowFlake(9)
}
func TestNewSnowFlake2(t *testing.T) {
	wa := new(sync.WaitGroup)
	for i := 0; i < 20; i++ {
		wa.Add(1)
		t := int64(i)
		go func(t int64) {
			defer wa.Done()
			snowFlake(t)
		}(t)
	}
	wa.Wait()
}
func BenchmarkSnowFlake_GetID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		snowFlake(10)
	}
}
func snowFlake(wrokid int64) {
	sf, err := NewSnowFlake(wrokid)
	if err != nil {
		fmt.Println(err)
	}
	ch := make(chan int64)
	mp := new(sync.Map)
	for i := 0; i < count; i++ {
		go func() {
			ch <- sf.GetID()
		}()
	}
L:
	for i := 0; i < count; i++ {
		select {
		case d := <-ch:
			if _, ok := mp.Load(d); ok {
				fmt.Println("bad", d)
				//break L
			}
			mp.Store(d,1)
		default:
			break L
		}
	}
}
