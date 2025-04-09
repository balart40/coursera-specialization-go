package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

The host allows no more than 2 philosophers to eat concurrently.

Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

Pases race

$ go run -race philosophers_problem.go
Kant starting to eat
Descartes starting to eat
Kant Eating
Descartes Eating
Descartes finish to eat meal 1
Kant finish to eat meal 1
Aristotle starting to eat
Aristotle Eating
Aristotle finish to eat meal 1
Socrates starting to eat
Socrates Eating
Socrates finish to eat meal 1
Plato starting to eat
Plato Eating
Aristotle starting to eat
Aristotle Eating
Aristotle finish to eat meal 2
Plato finish to eat meal 1
Descartes starting to eat
Descartes Eating
Descartes finish to eat meal 2
Socrates starting to eat
Socrates Eating
Socrates finish to eat meal 2
Kant starting to eat
Kant Eating
Kant finish to eat meal 2
Aristotle starting to eat
Aristotle Eating
Aristotle finish to eat meal 3
Plato starting to eat
Descartes starting to eat
Descartes Eating
Plato Eating
Plato finish to eat meal 2
Descartes finish to eat meal 3
Kant starting to eat
Socrates starting to eat
Kant Eating
Socrates Eating
Kant finish to eat meal 3
Socrates finish to eat meal 3
Plato starting to eat
Plato Eating
Plato finish to eat meal 3

*/

type ChopS struct{ sync.Mutex }

type Philo struct {
	name            string
	id              int
	leftCS, rightCS *ChopS
	eaten           int
	permitChan      chan struct{}
	doneChan        chan struct{}
}

type Host struct {
	chanBuff    int
	requestChan chan *Philo
}

func (h *Host) Authorize() {
	sem := make(chan struct{}, h.chanBuff) // Only 2 philosophers at a time

	for philo := range h.requestChan {
		sem <- struct{}{} // acquire slot (waits if 2 philosophers already eating)

		// Grant permission
		go func(p *Philo) {
			p.permitChan <- struct{}{} // notify philosopher to start

			// Wait until they're done
			<-p.doneChan

			<-sem // release slot
		}(philo)
	}
}

func (p *Philo) eat(host *Host, wg *sync.WaitGroup) {
	// We will eat 3 times
	for i := 0; i < 3; i++ {
		host.requestChan <- p // Ask for permission
		<-p.permitChan        // Wait until host says "go"

		// Lock chopsticks (order doesn't matter here)
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("finishing eating %d\n", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		p.doneChan <- struct{}{} // Notify host we're done
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	var host Host
	var channelBuffer int = 2
	var numberOfMeals int = 3
	var names []string = []string{"Plato", "Socrates", "Descartes",
		"Aristotle", "Kant"}
	host.chanBuff = channelBuffer
	host.requestChan = make(chan *Philo, channelBuffer)

	// Create 5 Chopsticks
	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	// Create Philosophers
	philos := make([]*Philo, 5)
	for i := 0; i < len(names); i++ {
		philos[i] = &Philo{
			name:       names[i],
			id:         i + 1,
			leftCS:     CSticks[i],
			rightCS:    CSticks[(i+1)%len(names)],
			eaten:      0,
			permitChan: make(chan struct{}),
			doneChan:   make(chan struct{}),
		}
	}

	//eat
	wg.Add(numberOfMeals * len(philos))

	go host.Authorize()

	for i := 0; i < len(philos); i++ {
		go philos[i].eat(&host, &wg)
	}

	wg.Wait()
}
