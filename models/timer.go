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
	return nil
}

func ( t Timer ) Update( msg tea.Msg ) ( Timer ,tea.Cmd ) {
	var cmd tea.Cmd
	switch msg :=  msg.( type )  {
		case TimerStartMsg:
			if !t.running {
				t.running = true 
				t.timer = timer.New( t.duration )
				cmd = t.Init()
			}
		case TimerStopMsg:
			if t.running {
				t.running = false
			}
		case TimerRestarttMsg:
			if t.running {
				t.running = false
			}
			t.timer = timer.New( t.duration )
		case TimerTickMsg:
			if t.running {
				var timerCmd tea.Cmd
				t.timer, timerCmd = t.timer.Update( msg )
				cmd = tea.Batch( cmd, timerCmd )
				cmd = tea.Batch( cmd, func() tea.Msg { 
					return TimerTickMsg{ TimeRemaining: t.timer.Timeout - time.Since( t.timer.StartTime ) }
				} )
			}
	}
	return t, cmd
}




