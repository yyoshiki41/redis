package internal

import (
	"time"
)

type Ticker struct {
	ticker *time.Ticker
	done   chan struct{}
}

func NewTicker(d time.Duration) *Ticker {
	t := &Ticker{
		ticker: time.NewTicker(d),
		done:   make(chan struct{}),
	}
	return t
}

func (t *Ticker) Stop() {
	t.ticker.Stop()
	close(t.done)
}

func (t *Ticker) Receive() <-chan time.Time {
	return t.ticker.C
}

func (t *Ticker) Done() <-chan struct{} {
	return t.done
}
