package ex9_5

import (
	"fmt"
	"time"
)

func PingPong() {
	ping := make(chan int)
	pong := make(chan int)
	done := make(chan bool)
	go func() {
		round := 0
		start := time.Now()
		seconds := 0.0
		for {
			ball := <-ping
			pong <- ball

			round++
			taken := time.Since(start)
			takenSeconds := taken.Seconds()
			if takenSeconds > seconds+5 {
				seconds = takenSeconds
				fmt.Printf("Round: %d, Taken: %v, Rate: %v rounds/s\n", round, taken, float64(round)/taken.Seconds())
			}
		}
	}()
	go func() {
		for {
			ball := <-pong
			ping <- ball
		}
	}()

	ping <- 99
	<-done
}
