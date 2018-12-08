package throttle

import "time"

type Throttle struct{
  PeriodicityMs int
	Limit int
	mwc *MovingWindowCounter
}

func( throttle *Throttle ) Invoke( f func() ){
	if throttle.mwc == nil {
		throttle.mwc = &MovingWindowCounter{PeriodicityMs: throttle.PeriodicityMs}
		throttle.mwc.Start()
	}
  for throttle.mwc.Count() > throttle.Limit {
		time.Sleep(2 * time.Millisecond)
	}
	throttle.mwc.Tick()
	f()
}
