package main

import "fmt"

const SECONDS_PER_MINUTE = 60
const MINUTES_PER_HOUR = 60
const SECONDS_PER_HOUR = MINUTES_PER_HOUR * SECONDS_PER_MINUTE

func main() {
	var timeInSeconds int
	fmt.Scanf("%d", &timeInSeconds)

	hours := timeInSeconds / SECONDS_PER_HOUR
	timeInSeconds -= hours * SECONDS_PER_HOUR
	minutes := timeInSeconds / SECONDS_PER_MINUTE
	timeInSeconds -= minutes * SECONDS_PER_MINUTE

	fmt.Printf("%0.2d:%0.2d:%0.2d\n", hours, minutes, timeInSeconds)
}
