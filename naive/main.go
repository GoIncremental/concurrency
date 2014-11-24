package main

import (
	"fmt"
	"math/rand"
	"time"
)

type die struct {
	value int
	id    int
}

func (d *die) printInfo() {

	fmt.Printf("Dice number %d is currently %d \n", d.id, d.value)

	//ch <- dice

}

type dice struct {
	phase int
	box   []*die
}

func (b *dice) printInfo() {
	fmt.Printf("The current phase is: %d\n", b.phase)

	for _, d := range b.box {
		d.printInfo()
	}

}

func orderDice(val int, d dice, ch chan dice) {

	d.phase = val
	for _, v := range d.box {

		v.value = val
		time.Sleep(time.Duration(100) * time.Millisecond)
	}

	ch <- d

}

func main() {

	start := time.Now()

	var box dice
	for e := 0; e < 10; e++ {

		d := &die{
			value: rand.Intn(5) + 1,
			id:    e + 1,
		}

		box.box = append(box.box, d)

	}
	ch := make(chan dice)

	for i := 0; i < 6; i++ {
		go orderDice(i+1, box, ch)

	}
	for g := 0; g < 6; g++ {
		output := <-ch
		output.printInfo()

	}
	timeTaken := time.Since(start)
	fmt.Printf("Time taken : %s \n", timeTaken)
}
