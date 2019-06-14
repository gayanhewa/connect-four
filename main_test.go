package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/gayanhewa/connect-four/interactions"
)

func TestInitPlayers(t *testing.T) {
	stdout := new(bytes.Buffer)
	stdin := strings.NewReader("G")
	i := interactions.NewInteractor(stdout, stdin)
	p1, p2 := initPlayers(i, 1)
	if reflect.TypeOf(p1) != reflect.TypeOf(p2) {
		t.Fatal("Failed incorrect player combination created")
	}
	p3, p4 := initPlayers(i, 2)
	if p3 == p4 {
		t.Fatal("Failed incorrect player combination created")
	}
	p5, p6 := initPlayers(i, 3)
	if reflect.TypeOf(p5) != reflect.TypeOf(p6) {
		t.Fatal("Failed incorrect player combination created")
	}
}
