// models/AppModel.go

package models

import (
	"gotype/conf"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AppModel struct {
	width		int
	height		int
	altscreenActive bool
	StatusBar	StatusBarModel
	Menu		MenuModel
	Game 		GameModel
	CurrentView	string
	config 		conf.Config
}

func NewAppModel() AppModel {

	c := conf.NewConfig()

	return AppModel{
		config: c,
		CurrentView: "MENU",
		StatusBar: NewStatusBarModel( c ),
		Menu: NewMenuModel( c ),
		Game: NewGameModel( c ),
	}
}

func ( m AppModel ) Init() tea.Cmd {
	return nil
}

func ( m AppModel ) Update( msg tea.Msg ) ( tea.Model, tea.Cmd ) {
	var cmd tea.Cmd
    var cmds []tea.Cmd
	
	switch msg := msg.( type ) {
		case tea.KeyMsg:
			if msg.String() == "ctrl+c" {
				return m, tea.Quit
			}
			newStatusBar, statusbarCmd := m.StatusBar.Update( msg )	
			m.StatusBar = newStatusBar.( StatusBarModel )
			cmd = statusbarCmd

		case EditorFinishedMsg:
			if msg.err != nil {
				m.CurrentView = "MENU"
				m.StatusBar.CurrentView = m.CurrentView
				m.altscreenActive = false
				cmd := tea.ExitAltScreen
				return m, cmd
			}

		case tea.WindowSizeMsg:
			m = m.handleResize( msg.Height, msg.Width )
		
		case ChangeViewMsg:
			m.CurrentView = msg.View
			m.StatusBar.CurrentView = msg.View
			return m, nil
		
	}
	switch( m.CurrentView ) {
		case "MENU":
			newMenu, menuCmd := m.Menu.Update( msg )
			m.Menu = newMenu.( MenuModel )
			cmd = menuCmd
				
		case "Game":
			newGame, gameCmd := m.Game.Update( msg )
			m.Game = newGame.( GameModel )
			cmd = gameCmd
						
	}

	
	cmds = append(cmds, cmd)
    return m, tea.Batch(cmds...)
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
		case "Game":
			mainView = lipgloss.JoinVertical( 0, mainWin.Render(), m.Game.View(), m.StatusBar.View() )
	}
	
	return mainView

}

func ( m AppModel ) handleResize( height, width int ) AppModel {
	
	m.height = height
	m.width = width
	
	m.StatusBar = m.StatusBar.handleResize( height, width )
	m.Menu = m.Menu.handleResize( height, width )
	m.Game = m.Game.handleResize( height, width )
	
	return m
}




