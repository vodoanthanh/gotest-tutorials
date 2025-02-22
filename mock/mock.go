package mock

import (
	"fmt"
	"time"
)

var doubleFunc = double

func double(val int) int {
	return val * 2
}

func doubleNumber(input int) {
	time.Sleep(1 * time.Second) // Simulate work
	fmt.Println("F1", doubleFunc(input))
}
