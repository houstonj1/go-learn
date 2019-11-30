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

// Sleep function
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Countdown function
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
