package main

import (
	"fmt"
	"testing"
)

func TestCall(t *testing.T) {
	var 本金 float64
	本金 = 20000
	for i := 0; i < 48; i++ {
		本金 = 本金*1.1 + 3000
	}
	sprintf := fmt.Sprintf("%f", 本金)
	fmt.Println(sprintf)
}
