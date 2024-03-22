package main

import (
	"github.com/xsni1/morse-code/state_machine"
)

type MorseDecoder struct {
	sm    state_machine.StateMachine
	state string
}

func NewMorseDecoder() MorseDecoder {
	decoder := MorseDecoder{}
	emptyState := state_machine.State{
		Next: map[string]state_machine.State{},
	}
	stateMachine := state_machine.StateMachine{
		Current: emptyState,
	}
	decoder.sm = stateMachine
	return decoder
}

func (d *MorseDecoder) Decode(encodedMsg string) string {
	msg := ""

	for _, char := range encodedMsg {
		d.sm.Process(string(char))
	}

	return msg
}

func main() {
	decoder := NewMorseDecoder()
	decoder.Decode("")
}
