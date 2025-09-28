// models/GameModel.go

package models

import (
	"fmt"
	"gotype/conf"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type GameModel struct {
	height		int
	width 		int
	textBuffer	[]string
	config		conf.Config
	targetText	[]string
	cursor		int
}

func ( m GameModel ) Init() tea.Cmd {
	return nil
}

func NewGameModel( c conf.Config ) GameModel { 
	return GameModel{
		config: c,
		cursor: 0,
		targetText: strings.Split( "testing a string" ,"" ),
		textBuffer: make([]string, 0),
	}
}


func ( m GameModel ) Update( msg tea.Msg ) ( tea.Model, tea.Cmd ) {

	switch msg := msg.( type ) {
		case tea.KeyMsg:
			switch( msg.String() ) { 
				case "esc":
					return m, ChangeView( "MENU" )

				case "backspace", "left":
					if m.cursor > 0 {
						m.cursor --
					}

				default: 
					m.textBuffer = append( m.textBuffer, msg.String() )
					if m.textBuffer[ m.cursor ] == m.targetText[ m.cursor ] && 
					m.cursor < len( m.targetText ) { 
						m.cursor ++
					}	
			}
	}
		return m, nil
}

func ( m GameModel ) View() string { 
	
	keystrokes := lipgloss.NewStyle().
		Align( lipgloss.Center )

	content := keystrokes.Render( fmt.Sprintf( " %s ", strings.Join( m.targetText, "" ) ) )
	/*
	for i, text := range m.targetText { 
		if m.cursor == i {
			cursorText := lipgloss.NewStyle().
				Foreground( m.config.Colors.Background ).
				Background( m.config.Colors.Foreground )
		}

		if m.cursor > i {	
			correctText := lipgloss.NewStyle().
				Foreground( m.config.Colors.Foreground ).
				Background( m.config.Colors.Background )
		}
		
		if m.cursor < i {
			untypedText := lipgloss.NewStyle(). 
				Foreground( m.config.Colors.Color7 ).
				Background( m.config.Colors.Background )
		}
		
	}	
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		
	) 
	*/
	return content 

}



func ( m GameModel ) handleResize( height, width int ) GameModel {
	
	m.height = height
	m.width = width
	return m
}
