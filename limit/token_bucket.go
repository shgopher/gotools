package limit

import (
	"time"
)

type TokenBucket struct {
	tokenPool chan bool
}

func NewTokenBucket(maxCount int64,interval time.Duration)*TokenBucket {
	t := TokenBucket{}
	t.tokenPool =  make(chan bool,maxCount)
	go func() {
		ti := time.NewTicker(time.Duration(interval.Nanoseconds()/(maxCount*1000*1000))* time.Millisecond)
		for  {
			select {
			case <- ti.C:
				<- t.tokenPool
			}
		}
	}()
	return &t
}
func(t *TokenBucket)Run(){
	t.tokenPool<- true// when the chan <- true , the length = length - 1
}
