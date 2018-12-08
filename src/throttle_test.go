package throttle 

import "testing"

func TestThrottle( t *testing.T ){
	throttle := &Throttle{PeriodicityMs: 100, Limit: 100}
	start := makeTimestamp()
	for i := 0; i < 200; i++ {
		throttle.Invoke(func() { _ = 1+1 })
	}
	end := makeTimestamp()
	if end-start < 200 {
		t.Error("expected to take at least 200ms but it happened faster:", end-start)
	}
}
