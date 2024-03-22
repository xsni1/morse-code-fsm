package state_machine

import "fmt"

type State struct {
	Val          string
	Next         map[string]*State
	OnTransition func(State)
}

type StateMachine struct {
	Current *State
}

func (sm *StateMachine) Process(input string) {
	fmt.Println("processing", input)
	if next, ok := sm.Current.Next[input]; ok {
		fmt.Println("current", sm.Current.Val, "next", next.Val)
        next.OnTransition(*sm.Current)
		sm.Current = next
	}
}
