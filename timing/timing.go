package timing

import (
	"time"
	"fmt"
)

const (
	bias         = 21 * time.Microsecond
	duration30Hz = (time.Second / 30) + bias
)

var master_ch chan chan time.Time

func init() {
	tick := time.NewTicker(duration30Hz)
	master_ch = make(chan chan time.Time)
	chans := make([]chan time.Time, 0)
	go func() {
		for {
			select {
			case t := <-tick.C:
				if (len(chans) == 0) { break }
				for _, c := range chans {
					c <- t // TODO: Remove chan that's closed
				}
			case ch := <-master_ch:
				chans = append(chans, ch)
			}
		}
	}()
}

func Subscribe() (ch chan time.Time) {
	ch = make(chan time.Time)
	master_ch <- ch
	return ch
}
