// models/GameModel.go

package models

import (
	"fmt"
	"gotype/conf"
	"gotype/database"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/timer"
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
	timeout		float64
	elapsed		float64
	running		bool
	playthrough	database.Playthrough
	mistakes	[]int
	
}

func ( m GameModel ) Init() tea.Cmd {
	return nil
}

func NewGameModel( c conf.Config ) GameModel { 
	return GameModel{
		config: c,
		cursor: 0,
		targetText: strings.Split( "#include <iostream>\nusing namespace std;\n\nint main() {\n\tcout << \"Hello World!\";\n\treturn 0;\n}" ,"" ),
		textBuffer: make( []string, 0 ),
		running: false,
	}
}

func ( m GameModel ) Update( msg tea.Msg ) ( tea.Model, tea.Cmd ) {

	switch msg := msg.( type ) {

		case tea.KeyMsg:
			switch( msg.String() ) { 
				case "esc":
					return m, StopGame()

				case "backspace", "left":
					if m.cursor > 0 {
						m.cursor --
					}

				default:
					keyMatch := false
					if msg.String() == "tab" && m.targetText[ m.cursor ] == "\t" {
						keyMatch = true
					}

					if msg.String() == "enter" && m.targetText[ m.cursor ] == "\n" {
						keyMatch = true
					}
					if msg.String() == m.targetText[ m.cursor ] {
						keyMatch = true
					}
					
					if keyMatch && m.cursor < len( m.targetText ) { 
						m.textBuffer = append( m.textBuffer, msg.String() )
						m.cursor ++
						if m.cursor >= len( m.targetText ) {
							return m, StopGame()
						} 
					} else {
						m.mistakes = append( m.mistakes, m.cursor )
					}


			}
			if !m.running {
				m.running = true
				return m, StartGame()
			}
			m.calculateStats()
		case timer.TimeoutMsg:
			return m, StopGame()
	}
		return m, nil
}

func ( m GameModel ) View() string { 
	
	//keystrokes := lipgloss.NewStyle().
	//	Align( lipgloss.Center )

	//content := keystrokes.Render( fmt.Sprintf( " %s ", strings.Join( m.targetText, "" ) ) )
	textbox := lipgloss.NewStyle().
		Width( m.width ).
		Height( m.height ).
		Align( lipgloss.Center, lipgloss.Center )
	gamebox := lipgloss.NewStyle().Align( lipgloss.Left, lipgloss.Center ).BorderStyle( lipgloss.NormalBorder() ).Padding( 1, 2 )

	cursorText := lipgloss.NewStyle().
		Foreground( m.config.Colors.Background ).
		Background( m.config.Colors.Foreground )

	correctText := lipgloss.NewStyle().
		Foreground( m.config.Colors.Foreground ).
		Background( m.config.Colors.Background )
	
	mistakeText := lipgloss.NewStyle().
		Foreground( m.config.Colors.Color9 ).
		Background( m.config.Colors.Background )

	untypedText := lipgloss.NewStyle(). 
		Foreground( m.config.Colors.Color5 ).
		Background( m.config.Colors.Background )

	var buidler strings.Builder

	for i, text := range m.targetText { 
		if m.cursor == i {
			buidler.WriteString( cursorText.Render( text ) )
		}

		if m.cursor > i {
			if slices.Contains( m.mistakes, i ) {
				buidler.WriteString( mistakeText.Render( text ) )
			} else {
				buidler.WriteString( correctText.Render( text ) )
			}
		}
		
		if m.cursor < i {
			buidler.WriteString( untypedText.Render( text ) )
		}
	}
	
	buidler.WriteString( fmt.Sprintf( "\n\ndebug\n-----\nw: %d h: %d\nelapsed: %.2f\ntime left: %.2f\nwpm: %.2f\ncpm: %.2f\npercentfinished: %.2f%%\nmistakes: %d\naccuracy: %.2f\nconsistency: %.2f", m.width, m.height, m.playthrough.ElapsedTime, m.timeout, m.playthrough.Wpm, m.playthrough.Cpm, m.playthrough.CompPercent, m.playthrough.Mistakes, m.playthrough.Accuracy, m.playthrough.Consistency ) )

	content := textbox.Render( gamebox.Render( buidler.String() ) )	

	return content

}

func ( m *GameModel ) handleResize( height, width int ) {
	m.height = height
	m.width = width
}

func ( m GameModel ) updateTimeout( timeout float64 ) GameModel {

	m.timeout = timeout / 1000000000
	m.playthrough.ElapsedTime = 30 - m.timeout
	return m
}

func ( m *GameModel ) calculateStats(  ) {
	
	chars := float64( m.cursor )
	words := chars / 5.0
	secs := m.playthrough.ElapsedTime
	m.playthrough.Cpm = ( chars / secs ) * 60.0
	m.playthrough.Wpm = ( words / secs ) * 60.0
	m.playthrough.CompPercent = chars * 100.00 / float64( len( m.targetText ) ) 
	m.playthrough.Accuracy = float64( len( m.mistakes ) ) / float64( m.cursor )
	m.playthrough.Mistakes = len( m.mistakes )
}

func ( m *GameModel ) SaveGame(  ) {
	m.playthrough.Save()	
}

