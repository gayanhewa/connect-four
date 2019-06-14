package interactions

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrompt(t *testing.T) {
	stdout := new(bytes.Buffer)
	stdin := strings.NewReader("answer")
	i := NewInteractor(stdout, stdin)
	answer := i.Prompt("Question")
	if answer != "answer" {
		t.Fatalf("Failed asserting %q is equal to %q", answer, "answer")
	}
	if stdout.String() != "Question" {
		t.Fatalf("Failed asserting %q is equal to %q", "Question", stdout.String())
	}
}

func TestAnswer(t *testing.T) {
	stdout := new(bytes.Buffer)
	stdin := strings.NewReader("answer")
	i := NewInteractor(stdout, stdin)
	i.Prompt("Question")
	if i.Answer() != "answer" {
		t.Fatalf("Failed asserting %q is equal to %q", i.Answer(), "answer")
	}
}

func TestNotify(t *testing.T) {
	stdout := new(bytes.Buffer)
	stdin := strings.NewReader("")
	i := NewInteractor(stdout, stdin)
	i.Notify("message")
	if stdout.String() != "message" {
		t.Fatalf("Failed asserting %q is equal to %q", stdout.String(), "message")
	}
}
