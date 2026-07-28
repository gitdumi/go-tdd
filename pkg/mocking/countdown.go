package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

// can be deleted

type ConfigurableSleeper struct {
	Duration time.Duration
	CustomSleep    func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.CustomSleep(s.Duration)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}

	fmt.Fprint(out, finalWord)
}
