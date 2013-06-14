package timing

import (
	"time"
	"testing"
	"math"
	"./../util"
)

const (
	tol = 0.1
	target = 30
)

func TestTiming(t *testing.T) {
	// TODO: Add multiple goroutines subscribed to chan and test timing of each
	ch := Subscribe()
	timeout := time.After(1 * time.Second)
	t0 := time.Now()
	t1 := time.Now()
	measure := make([]float64, 0)
	for i := 0; i <= 10; i++ {
		select {
		case <-timeout:
			t.Fatal("Test Timeout")
			return
		case <-ch:
			t1 = time.Now()
			if i == 0 {
				t0 = time.Now()
				break
			}
			measure = append(measure, float64(t1.Sub(t0)))
			t0 = time.Now()
		}
	}
	hz_measure := float64(time.Second) / util.Average(measure)
	if (math.Abs(hz_measure - target) > tol) { 
		t.Fatal("\n\nTiming Failed: ", 
			hz_measure,
			"Hz\nTarget: ",
			target,
			"Hz\nTol: ",
			tol,
			"Hz\n")
	}
	return
}
