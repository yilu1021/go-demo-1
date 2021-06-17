package mascot_test

import (
	"testing"

	"example.com/go-demo-1/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMoscot() != "Go Gopher" {
		t.Fatal("Wrong mascot")
	}
}
