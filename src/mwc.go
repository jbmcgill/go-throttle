package throttle

import "sync"
import "time"
import "math"

type CountBuffer struct {
	sync.RWMutex
	frame   [2]int
	updated int64
}

type MovingWindowCounter struct {
	PeriodicityMs int
	counts        CountBuffer
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (mwc *MovingWindowCounter) Tick() {
	mwc.counts.Lock()
	defer mwc.counts.Unlock()
	mwc.counts.frame[1]++
}

func (mwc *MovingWindowCounter) Count() int {
	mwc.counts.Lock()
	defer mwc.counts.Unlock()
	offset := makeTimestamp() - mwc.counts.updated
	ratio := 1 - float64(offset)/1000
	return mwc.counts.frame[1] + int(math.Round(float64(mwc.counts.frame[0])*ratio))
}

func (mwc *MovingWindowCounter) Start() {
	mwc.counts.Lock()
	mwc.counts.updated = makeTimestamp()
	mwc.counts.Unlock()
	go func(mwc *MovingWindowCounter) {
		timer := time.Tick(time.Duration(mwc.PeriodicityMs) * time.Millisecond)
		for _ = range timer {
			mwc.counts.Lock()
			mwc.counts.frame[0] = mwc.counts.frame[1]
			mwc.counts.frame[1] = 0
			mwc.counts.updated = makeTimestamp()
			mwc.counts.Unlock()
		}
	}(mwc)
}
