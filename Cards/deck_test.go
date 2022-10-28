package main

import (
	"fmt"
	"testing"
)

func testNewDeck(t *testing.T) {
	d := newDeck()
	fmt.Println(d)
	if len(d) != 12 {
		t.Errorf("Expected length is 12")
	}
}
