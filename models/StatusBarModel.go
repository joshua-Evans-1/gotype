// models/StatusBarModel.go

package models

import (
	"fmt"
	"strings"
	"gotype/conf"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StatusBarModel struct {
	height		int
	width		int
	KeyBuffer	[]string
	CurrentView	string
	config 		conf.Config
}

func NewStatusBarModel( c conf.Config ) StatusBarModel {
	return StatusBarModel{
		config: c,
		CurrentView: "MENU",
	}
}

func ( m StatusBarModel ) Init() tea.Cmd {
	return nil
}

func ( m StatusBarModel ) Update( msg tea.Msg ) ( tea.Model, tea.Cmd ) {

	switch msg := msg.( type ) {
		case tea.KeyMsg:
			m.KeyBuffer = append( m.KeyBuffer, msg.String() )
				if len( m.KeyBuffer ) > 5 {
						m.KeyBuffer = m.KeyBuffer[1:]
				}
	}
	return m, nil
}

func ( m StatusBarModel ) View() string {
	
	mode := lipgloss.NewStyle().
		Background( m.config.Colors.Foreground ).
		Foreground( m.config.Colors.Background )

	keystrokes := lipgloss.NewStyle().
		Align( lipgloss.Right )

	leftContent := mode.Render( fmt.Sprintf( " %s ", m.CurrentView ) ) 
	rightContent := keystrokes.Render( fmt.Sprintf( " %s ", strings.Join( m.KeyBuffer, "" ) ) )
	spacer := lipgloss.NewStyle().
		Width( m.width - lipgloss.Width( rightContent ) - lipgloss.Width( leftContent ) ).
		Render()
	
	statusBar :=
		lipgloss.JoinHorizontal(
			0,
			leftContent,
			spacer,
			rightContent,
		)
		
	return statusBar
}

func ( m StatusBarModel ) handleResize( height, width int ) StatusBarModel {
	
	m.height = height
	m.width = width
	return m
}
