package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper type
type Sleeper interface {
	Sleep()
}

// DefaultSleeper struct
type DefaultSleeper struct{}

// ConfigurableSleeper struct
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep function
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (c *ConfigurableSleeper) Sleep() {
}

// Countdown function
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
