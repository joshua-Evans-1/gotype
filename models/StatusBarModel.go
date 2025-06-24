// models/StatusBarModel.go

package models

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)



type StatusBarModel struct {
	height		int
	width		int
	KeyBuffer	[]string
	CurrentView	string
}

func ( m StatusBarModel ) Init() tea.Cmd {
	return nil
}



func ( m StatusBarModel ) Update( msg tea.Msg ) ( tea.Model, tea.Cmd ) {
	return m, nil
}

func ( m StatusBarModel ) View() string {
	
	mode := lipgloss.NewStyle().
		Background( lipgloss.Color( "#7D56F4" ) ).
		Foreground( lipgloss.Color( "#EAEBD0" ) )

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



