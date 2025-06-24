// models/AppModel.go

package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"os/exec"
	"gotype/conf"
)

type AppModel struct {
	width		int
	height		int
	altscreenActive bool
	StatusBar	StatusBarModel
	Menu		MenuModel
	CurrentView	string
	config 		conf.Config

	//levels 		[]level

}

func NewAppModel() AppModel {
	appModel := AppModel{}

	appModel.CurrentView = "MENU"
	appModel.StatusBar.CurrentView = appModel.CurrentView
	appModel.Menu.menuOptions = []string{ "Levels", "Settings", "HELP" }
	appModel.Menu.cursor = 0
	return appModel 
}

func ( m AppModel ) Init() tea.Cmd {
	return nil
}

func ( m AppModel ) Update( msg tea.Msg ) ( tea.Model, tea.Cmd ) {
	switch msg := msg.( type ) {
		case tea.KeyMsg:
			if msg.String() == "ctrl+c" {
				return m, tea.Quit
			}

			m.StatusBar.KeyBuffer = append( m.StatusBar.KeyBuffer, msg.String() )
			if len( m.StatusBar.KeyBuffer ) > 5 {
				m.StatusBar.KeyBuffer = m.StatusBar.KeyBuffer[1:]
			}
			switch( m.CurrentView ) {
				case "MENU":
					switch( msg.String() ) {
						case "j", "down":
							if m.Menu.cursor >= 2 {
								m.Menu.cursor = 0
							} else {
								m.Menu.cursor++
							}
						case "k", "up":
							if m.Menu.cursor <= 0 {
								m.Menu.cursor = 2
							} else {
								m.Menu.cursor--
							}
						case "enter", "l", "right":
							if m.Menu.menuOptions[ m.Menu.cursor ] == "Settings" {
								m.altscreenActive = true
								openEditor()
								//cmd := tea.EnterAltScreen
								return m, openEditor()							} else {
								m.CurrentView = m.Menu.menuOptions[ m.Menu.cursor ]
								m.StatusBar.CurrentView = m.CurrentView
							}
						case "?":
							m.CurrentView = "HELP"
							m.StatusBar.CurrentView = m.CurrentView
						case "q", "esc":
							return m, tea.Quit
					}
				case "HELP":
					switch( msg.String() ) {
						case "j", "down":
						case "k", "up":
						case "q", "esc":
							m.CurrentView = "MENU"
							m.StatusBar.CurrentView = m.CurrentView
					}
				
			}
		case editorFinishedMsg:
			if msg.err != nil {
				m.CurrentView = "MENU"
				m.StatusBar.CurrentView = m.CurrentView
				m.altscreenActive = false
				cmd := tea.ExitAltScreen
				return m, cmd
			}
		case tea.WindowSizeMsg:
			m.Menu.width = msg.Width
			m.Menu.height = msg.Height
			m.StatusBar.width = msg.Width
			m.StatusBar.height = msg.Height
			m.width = msg.Width
			m.height = msg.Height
	}
	return m, nil
}

func ( m AppModel ) View() string {

	mainWin := lipgloss.NewStyle().
		Height( m.height - 1 ).
		Width( m.width - 2 )
	mainView := ""
	switch( m.CurrentView ) {
		case "MENU":
			mainView = lipgloss.JoinVertical( 0, mainWin.Render( m.Menu.View() ), m.StatusBar.View() )

		case "HELP":
			mainView = lipgloss.JoinVertical( 0, mainWin.Render(  ), m.StatusBar.View() )
	}
	
	return mainView

}

type editorFinishedMsg struct{ err error }

func openEditor() tea.Cmd {
	editor := os.Getenv( "EDITOR" )
	home := os.Getenv( "HOME" )
	if editor == "" {
		editor = "vim"
	}
	c := exec.Command( editor, home + "/.config/gotype/conf" ) 
	return tea.ExecProcess( c, func( err error ) tea.Msg {
		return editorFinishedMsg{ err }
	})
}
