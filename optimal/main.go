package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type die struct {
	value int
	id    int
}

type person struct {
	id  int
	in  chan *die
	out chan *die
}

func (p *person) doWork() {

	for die := range p.in {
		if p.id == 1 {
			die.value = 1
		} else {
			die.value++
		}
		log.Printf("Person %d set die %d to %d\n", p.id, die.id, die.value)
		time.Sleep(time.Duration(100) * time.Millisecond)
		p.out <- die
	}

}

func logCompletion(c chan *die) {

	d := <-c
	log.Printf("Die %d is completed\n", d.id)

}

func main() {

	start := time.Now()

	//create 10 dice
	var dice []*die
	for e := 0; e < 10; e++ {

		d := &die{
			value: rand.Intn(5) + 1,
			id:    e + 1,
		}

		dice = append(dice, d)

	}

	//Create 6 people
	var room []*person
	startCh := make(chan *die, 10)
	prevCh := startCh
	for i := 0; i < 6; i++ {
		out := make(chan *die, 10)
		p := &person{
			id:  i + 1,
			in:  prevCh,
			out: out,
		}

		prevCh = out

		room = append(room, p)
		go p.doWork()

	}

	for _, d := range dice {

		startCh <- d

	}

	for _ = range dice {

		logCompletion(prevCh)

	}

	timeTaken := time.Since(start)
	fmt.Printf("Time taken : %s \n", timeTaken)

}
