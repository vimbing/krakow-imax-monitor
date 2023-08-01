package main

import (
	"cinemacity/internal/monitor"
	"fmt"
)

func main() {
	monitor, err := monitor.Init()

	if err != nil {
		fmt.Println(err.Error())
	}

	monitor.Monitor()
}
