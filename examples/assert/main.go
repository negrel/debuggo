package main

import (
	"fmt"

	"github.com/negrel/debuggo/pkg/assert"
)

func getWeather() (int, error) {
	return -1, fmt.Errorf("unable to get the weather forecast")
}

func main() {
	weather, err := getWeather()
	assert.Nil(err)

	fmt.Println("In one hour the temperature will be", weather, "°C")
}
