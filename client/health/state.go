package health

import "sync/atomic"

type State struct {
	ok atomic.Bool
}

func NewState() *State {
	s := &State{}
	s.ok.Store(true)
	return s
}

func (s *State) SetHealthy()   { s.ok.Store(true) }
func (s *State) SetUnhealthy() { s.ok.Store(false) }
func (s *State) IsHealthy() bool {
	return s.ok.Load()
}
