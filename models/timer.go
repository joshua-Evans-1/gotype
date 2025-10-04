// models/timer.go

package models

import (
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type Timer struct {
	timer timer.Model
	running bool
	duration time.Duration
}


type TimerStartMsg struct{}
type TimerStopMsg struct{}
type TimerRestarttMsg struct{}
type TimerTickMsg struct{
	TimeRemaining time.Duration
}

func newTimer( duration time.Duration ) Timer {
	return Timer{
		timer: timer.New( duration ),
		running: false,
		duration: duration,
	}
}

func ( t Timer ) Init() tea.Cmd {
	return t.timer.Init(  )
}

func ( t Timer ) Update( msg tea.Msg ) ( Timer ,tea.Cmd ) {
	switch msg :=  msg.( type )  {
		case TimerTickMsg:
			var cmd tea.Cmd
			t.timer, cmd = t.timer.Update( msg )
			return t, cmd

		case timer.StartStopMsg:
			var cmd tea.Cmd
			t.timer, cmd = t.timer.Update( msg )
			return t, cmd

		case timer.TimeoutMsg:
			var cmd tea.Cmd
			t.running = false
			t.timer, cmd = t.timer.Update( msg )
			return t, cmd 

	}
	return t, nil
}




