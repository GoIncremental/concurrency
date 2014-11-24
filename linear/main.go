package main

import (
	"fmt"
	"math/rand"
	"time"
)

type die struct {
	value int
}

func orderDice(val int, dice []*die) {

	for i, v := range dice {
		fmt.Printf("Dice number %d is currently %d \n", i+1, v.value)
		fmt.Printf("Setting Dice %d to: %d \n", i+1, val)
		v.value = val
		time.Sleep(100 * time.Millisecond)
	}

}

func main() {

	start := time.Now()

	var dice []*die
	for e := 0; e < 10; e++ {

		d := &die{
			value: rand.Intn(5) + 1,
		}

		dice = append(dice, d)

	}
	for i := 0; i < 6; i++ {

		orderDice(i+1, dice)

	}

	timeTaken := time.Since(start)
	fmt.Printf("Time taken : %s \n", timeTaken)
}
