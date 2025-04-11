package elocalculator

import "testing"

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
	if writer.output != expected {
		t.Errorf(`expected: %q, actual: %q`, expected, writer.output)
	}
}
