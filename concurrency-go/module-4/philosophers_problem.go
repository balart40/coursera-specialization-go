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
}

type Host struct {
	chanBuff int
	channel  chan *Philo
}

func (host *Host) Authorise(wg *sync.WaitGroup) {
	for {
		if len(host.channel) == host.chanBuff {
			<-host.channel
			<-host.channel

			// Add delay
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func (p *Philo) eat(host *Host, wg *sync.WaitGroup) {
	// We will eat 3 times
	for i := 0; i < 3; i++ {
		// small fix since philosphers were eating 4 times
		if p.eaten < 3 {
			host.channel <- p
			p.leftCS.Lock()
			p.rightCS.Lock()

			fmt.Printf("%v starting to eat \n", p.name)
			p.eaten++
			fmt.Printf("%v finishing eating \n", p.name)
			
			p.rightCS.Unlock()
			p.leftCS.Unlock()
			wg.Done()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var host Host
	var channelBuffer int = 2
	var numberOfMeals int = 3
	var names []string = []string{"Plato", "Socrates", "Descartes", "Aristotle", "Kant"}
	host.chanBuff = channelBuffer
	host.channel = make(chan *Philo, channelBuffer)

	// Create 5 Chopsticks
	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	// Create Philosophers
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{names[i], i + 1, CSticks[i], CSticks[(i+1)%5], 0}
	}

	//eat
	wg.Add(numberOfMeals * len(philos))

	go host.Authorise(&wg)

	for i := 0; i < len(philos); i++ {
		go philos[i].eat(&host, &wg)
	}

	wg.Wait()
}
