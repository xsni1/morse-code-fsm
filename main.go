package main

import (
	"fmt"

	"github.com/xsni1/morse-code/state_machine"
)

type MorseDecoder struct {
	sm    state_machine.StateMachine
	state string
}

func NewMorseDecoder() *MorseDecoder {
	decoder := &MorseDecoder{}
	EMPTY_STATE := &state_machine.State{OnTransition: func(s state_machine.State) {
        fmt.Println("empty", s)
		switch s.Val {
		case "A":
			decoder.state += "A"
		case "B":
			decoder.state += "B"
		case "C":
			decoder.state += "C"
		case "D":
			decoder.state += "D"
		case "E":
			decoder.state += "E"
		case "F":
			decoder.state += "F"
		case "G":
			decoder.state += "G"
		case "H":
			decoder.state += "H"
		case "I":
			decoder.state += "I"
		case "J":
			decoder.state += "J"
		case "K":
			decoder.state += "K"
		case "L":
			decoder.state += "L"
		case "M":
			decoder.state += "M"
		case "N":
			decoder.state += "N"
		case "O":
			decoder.state += "O"
		case "P":
			decoder.state += "P"
		case "Q":
			decoder.state += "Q"
		case "R":
			decoder.state += "R"
		case "S":
			decoder.state += "S"
		case "T":
			decoder.state += "T"
		case "U":
			decoder.state += "U"
		case "V":
			decoder.state += "V"
		case "W":
			decoder.state += "W"
		case "X":
			decoder.state += "X"
		case "Y":
			decoder.state += "Y"
		case "Z":
			decoder.state += "Z"
		}
	}}
	E_STATE := &state_machine.State{Val: "E", OnTransition: func(s state_machine.State) {}}
	I_STATE := &state_machine.State{Val: "I", OnTransition: func(s state_machine.State) {}}
	S_STATE := &state_machine.State{Val: "S", OnTransition: func(s state_machine.State) {}}
	U_STATE := &state_machine.State{Val: "U", OnTransition: func(s state_machine.State) {}}
	F_STATE := &state_machine.State{Val: "F", OnTransition: func(s state_machine.State) {}}
	H_STATE := &state_machine.State{Val: "H", OnTransition: func(s state_machine.State) {}}
	V_STATE := &state_machine.State{Val: "V", OnTransition: func(s state_machine.State) {}}
	A_STATE := &state_machine.State{Val: "A", OnTransition: func(s state_machine.State) {}}
	W_STATE := &state_machine.State{Val: "W", OnTransition: func(s state_machine.State) {}}
	P_STATE := &state_machine.State{Val: "P", OnTransition: func(s state_machine.State) {}}
	J_STATE := &state_machine.State{Val: "J", OnTransition: func(s state_machine.State) {}}
	R_STATE := &state_machine.State{Val: "R", OnTransition: func(s state_machine.State) {}}
	L_STATE := &state_machine.State{Val: "L", OnTransition: func(s state_machine.State) {}}
	T_STATE := &state_machine.State{Val: "T", OnTransition: func(s state_machine.State) {}}
	N_STATE := &state_machine.State{Val: "N", OnTransition: func(s state_machine.State) {}}
	D_STATE := &state_machine.State{Val: "D", OnTransition: func(s state_machine.State) {}}
	B_STATE := &state_machine.State{Val: "B", OnTransition: func(s state_machine.State) {}}
	X_STATE := &state_machine.State{Val: "X", OnTransition: func(s state_machine.State) {}}
	K_STATE := &state_machine.State{Val: "K", OnTransition: func(s state_machine.State) {}}
	C_STATE := &state_machine.State{Val: "C", OnTransition: func(s state_machine.State) {}}
	Y_STATE := &state_machine.State{Val: "Y", OnTransition: func(s state_machine.State) {}}
	M_STATE := &state_machine.State{Val: "M", OnTransition: func(s state_machine.State) {}}
	O_STATE := &state_machine.State{Val: "O", OnTransition: func(s state_machine.State) {}}
	G_STATE := &state_machine.State{Val: "G", OnTransition: func(s state_machine.State) {}}
	Z_STATE := &state_machine.State{Val: "Z", OnTransition: func(s state_machine.State) {}}
	Q_STATE := &state_machine.State{Val: "Q", OnTransition: func(s state_machine.State) {}}

	EMPTY_STATE.Next = map[string]*state_machine.State{
		".": E_STATE,
		"_": T_STATE,
	}
	E_STATE.Next = map[string]*state_machine.State{
		".": I_STATE,
		"_": A_STATE,
		" ": EMPTY_STATE,
	}
	I_STATE.Next = map[string]*state_machine.State{
		".": S_STATE,
		"_": U_STATE,
		" ": EMPTY_STATE,
	}
	A_STATE.Next = map[string]*state_machine.State{
		".": R_STATE,
		"_": W_STATE,
		" ": EMPTY_STATE,
	}
	S_STATE.Next = map[string]*state_machine.State{
		".": H_STATE,
		"_": V_STATE,
		" ": EMPTY_STATE,
	}
	H_STATE.Next = map[string]*state_machine.State{
		" ": EMPTY_STATE,
	}
	V_STATE.Next = map[string]*state_machine.State{
		" ": EMPTY_STATE,
	}
	U_STATE.Next = map[string]*state_machine.State{
		".": F_STATE,
		" ": EMPTY_STATE,
	}
	F_STATE.Next = map[string]*state_machine.State{
		" ": EMPTY_STATE,
	}
	R_STATE.Next = map[string]*state_machine.State{
		".": L_STATE,
		" ": EMPTY_STATE,
	}
	L_STATE.Next = map[string]*state_machine.State{
		" ": EMPTY_STATE,
	}
	W_STATE.Next = map[string]*state_machine.State{
		".": P_STATE,
		"_": J_STATE,
		" ": EMPTY_STATE,
	}
	P_STATE.Next = map[string]*state_machine.State{
		" ": EMPTY_STATE,
	}
	J_STATE.Next = map[string]*state_machine.State{
		" ": EMPTY_STATE,
	}
	T_STATE.Next = map[string]*state_machine.State{
		".": N_STATE,
		"_": M_STATE,
		" ": EMPTY_STATE,
	}
	N_STATE.Next = map[string]*state_machine.State{
		".": D_STATE,
		"_": K_STATE,
		" ": EMPTY_STATE,
	}
	D_STATE.Next = map[string]*state_machine.State{
		".": B_STATE,
		"_": X_STATE,
		" ": EMPTY_STATE,
	}
	B_STATE.Next = map[string]*state_machine.State{
		" ": EMPTY_STATE,
	}
	X_STATE.Next = map[string]*state_machine.State{
		" ": EMPTY_STATE,
	}
	K_STATE.Next = map[string]*state_machine.State{
		".": C_STATE,
		"_": Y_STATE,
		" ": EMPTY_STATE,
	}
	C_STATE.Next = map[string]*state_machine.State{
		" ": EMPTY_STATE,
	}
	Y_STATE.Next = map[string]*state_machine.State{
		" ": EMPTY_STATE,
	}
	M_STATE.Next = map[string]*state_machine.State{
		".": G_STATE,
		"_": O_STATE,
		" ": EMPTY_STATE,
	}
	O_STATE.Next = map[string]*state_machine.State{
		" ": EMPTY_STATE,
	}
	G_STATE.Next = map[string]*state_machine.State{
		".": Z_STATE,
		"_": Q_STATE,
		" ": EMPTY_STATE,
	}
	Z_STATE.Next = map[string]*state_machine.State{
		" ": EMPTY_STATE,
	}
	Q_STATE.Next = map[string]*state_machine.State{
		" ": EMPTY_STATE,
	}

	stateMachine := state_machine.StateMachine{
		Current: EMPTY_STATE,
	}
	decoder.sm = stateMachine
	return decoder
}

func (d *MorseDecoder) Decode(encodedMsg string) string {
	for _, char := range encodedMsg {
		d.sm.Process(string(char))
	}

	return d.state
}

func main() {
	decoder := NewMorseDecoder()
    // work around here
	fmt.Println(decoder.Decode(".... . ._.. ._.. ___ "))
}
