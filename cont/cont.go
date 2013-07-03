package cont

import (
	"sync"
	"time"
)

const (
	numWorkers = 5
)

type Stepper interface {
	Step(wg sync.WaitGroup)
	SetStep(tm time.Duration)
}

var master_ch chan Stepper
var wg sync.WaitGroup

func init() {
	master_ch = make(chan Stepper)
	for i := 0; i < numWorkers; i++ {
		go func() {
			for {
				st := <-master_ch
				st.Step(wg)
			}
		}()
	}
}

type ContinuousStep struct {
	tm       time.Duration
	steppers []Stepper
	numSteps int
}

func NewContinuousStep(tm time.Duration, steppers []Stepper, numSteps int) *ContinuousStep {
	c := new(ContinuousStep)
	c.tm = tm
	c.steppers = steppers
	c.numSteps = numSteps
	return c
}

// StepThrough calls each function registered to the ContinuousStep struct,
// waits for them all to complete and call wg.Done(), then repeats.
func (c *ContinuousStep) StepThrough() {
	for _, st := range c.steppers {
		st.SetStep(c.tm)
	}
	for i := 0; i < c.numSteps; i++ {
		for _, st := range c.steppers {
			wg.Add(1)
			master_ch <- st
		}
		wg.Wait()
	}
}
