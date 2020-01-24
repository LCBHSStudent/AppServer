package common

import (
	"log"
	"testing"
)

func TestCrypt(t *testing.T) {
	log.Println("Uint crypt test")
	if err := Crypt(); err != nil {
		t.Fatal(err)
	} else {
		t.Log("\nPass")
	}
}
