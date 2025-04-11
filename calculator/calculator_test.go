package elocalculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StringOutputWriter struct {
	output string
}

func (writer *StringOutputWriter) PrintLine(line string) {
	writer.output = line
	writer.output += "\n"
}

func TestItDisplaysAMessage(t *testing.T) {
	writer := StringOutputWriter{""}

	EloCalculator(writer.PrintLine)

	expected := "Hello Elo Calculator\n"

	assert.Equal(t, expected, writer.output)
}
