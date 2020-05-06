package id

import (
	"fmt"
	"sync"
	"time"
)

const (
	MAX_NUMBER       = 10
	MAX_WORK         = 12
	MAX_NUMBER_VALUE = 0x3ff
	MAX_WORK_VALUE   = 0xfff
	OFFSET_TIME      = MAX_NUMBER + MAX_WORK
	OFFSET_WORK      = MAX_NUMBER
)

var (
	oldNow int64 = 1588760321035
)

type SnowFlake struct {
	sync.Mutex
	timeStamp int64
	workNode  int64
	number    int64
}

func NewSnowFlake(workNode int64) (*SnowFlake, error) {
	if workNode > MAX_WORK_VALUE || workNode < 0 {
		return nil, fmt.Errorf("input workNode too big or < 0 ,right one is [0, 2^12+2^11+...2^0]")
	}
	return &SnowFlake{workNode: workNode}, nil
}

func (s *SnowFlake) GetID() int64 {
	s.Lock()
	defer s.Unlock()
	now := time.Now().UnixNano() / 1e6
	if now == s.timeStamp {
		s.number++
		if s.number > MAX_NUMBER_VALUE {
			for now <= s.timeStamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.number = 0
		s.timeStamp = now
	}
	return (s.timeStamp-oldNow)<<OFFSET_TIME | s.workNode<<OFFSET_WORK | s.number
}
