package main

import "testing"

func TestSoma(t *testing.T) {
	if soma(5, 5) != 10 {
		t.Error("Soma incorreta")
	}
}
