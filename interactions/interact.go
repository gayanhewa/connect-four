package interactions

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Question interface.
type Question interface {
	Prompt(prompt string) string
	Answer() string
}

// Notifier pushes a message to stdout.
type Notifier interface {
	Notify(message string)
}

// Interactor composes both Question and Notifier.
type Interactor interface {
	Question
	Notifier
}

// Interaction for the terminal.
type Interaction struct {
	w      io.Writer
	r      io.Reader
	prompt string
	answer string
}

// NewInteractor factory.
func NewInteractor(w io.Writer, r io.Reader) Interactor {
	return &Interaction{
		w: w,
		r: r,
	}
}

// Prompt question it self.
func (i *Interaction) Prompt(prompt string) string {
	i.prompt = prompt
	reader := bufio.NewReader(i.r)
	fmt.Fprintf(i.w, i.prompt)
	answer, _ := reader.ReadString('\n')
	i.answer = strings.TrimSuffix(answer, "\n")
	i.prompt = prompt
	return i.answer
}

// Answer for the question.
func (i *Interaction) Answer() string {
	return i.answer
}

// Notify the stdout.
func (i *Interaction) Notify(message string) {
	fmt.Fprintf(i.w, message)
}
