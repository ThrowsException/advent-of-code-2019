package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func calcFuel(val int) int {
	total := int(math.Floor(float64(val/3))) - 2
	if total > 0 {
		return total
	}
	return 0
}

func calcFuelIterative(val int) int {
	total := calcFuel(val)

	additionalFuel := calcFuel(total)
	for additionalFuel > 0 {
		total += additionalFuel
		additionalFuel = calcFuel(additionalFuel)
	}

	return total
}

func calcFuelRecursive(val int) int {
	total := calcFuel(val)

	if total == 0 {
		return 0
	}
	return total + calcFuelRecursive(total)
}

func main() {
	total := 0
	f, err := os.Open("fuel.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		total += calcFuelIterative(val)
	}
	fmt.Println(total)
}
