package throttle

import "testing"
import "time"

func TestBaseFunctionality(t *testing.T) {
	mwc := &MovingWindowCounter{PeriodicityMs: 1000}
	mwc.Start()
	mwc.Tick()
	mwc.Tick()
	expected := 2
	result := mwc.Count()
	if result != expected {
		t.Error("expected: ", expected, "but got:", result)
	}
}

func TestMovingWindow(t *testing.T) {
	periodicity := 100
	mwc := &MovingWindowCounter{PeriodicityMs: periodicity}
	mwc.Start()
	mwc.Tick()
	mwc.Tick()
	mwc.Tick()
	mwc.Tick()
	result := mwc.Count()
	expected := 4
	if result != expected {
		t.Error("expected:", expected, "but got:", result)
	}
	time.Sleep(110 * time.Millisecond)
	mwc.Tick()
	result = mwc.Count()
	expected = 5
	if result != expected {
		t.Error("expected:", expected, "but got:", result)
	}
	time.Sleep(200 * time.Millisecond)
	result = mwc.Count()
	expected = 0
	if result != expected {
		t.Error("expected:", expected, "but got:", result)
	}
}
