package state_machine

type State struct {
	Next         map[string]State
	OnTransition func()
}

type StateMachine struct {
	Current State
}

func (sm *StateMachine) Process(input string) {
	if next, ok := sm.Current.Next[input]; ok {
		sm.Current.OnTransition()
		sm.Current = next
	}
}
