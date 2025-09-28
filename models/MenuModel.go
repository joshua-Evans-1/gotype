// models/MenuModel.go

package models

import (
	"strings"
	"gotype/conf"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuModel struct {
	width		int
	height		int
	menuOptions	[]string
	cursor		int
	selected 	string
	config 		conf.Config
}

func NewMenuModel( c conf.Config ) MenuModel { 
	return MenuModel{
		config: c,
		cursor: 0,
		menuOptions: []string{ "Game", "Levels", "Settings", "Help" },
	}	
}


func ( m MenuModel ) Init() tea.Cmd {
	return nil
}

func ( m MenuModel ) Update( msg tea.Msg ) ( tea.Model, tea.Cmd ) {
	switch msg := msg.( type ) {
		case tea.KeyMsg:
			switch( msg.String() ) {
				case "j", "down":
					if m.cursor >= len( m.menuOptions ) {
						m.cursor = 0
					} else {
						m.cursor++
					}
				case "k", "up":
					if m.cursor <= 0 {
						m.cursor = len( m.menuOptions )
					} else {
						m.cursor--
					}
				case "enter", "l", "right":
					if m.menuOptions[ m.cursor ] == "Settings" {
						return m, openEditor()	
					} else {
						return m, ChangeView( m.menuOptions[ m.cursor ] )
					}
				case "?":
					return m, ChangeView("Help")
				case "q", "esc":
					return m, tea.Quit
			}
	}
	return m, nil
}


func (m MenuModel) View() string {
    menuWin := lipgloss.NewStyle().
        Width(m.width).
        MaxWidth(m.width).
        Align(lipgloss.Center)

    titleText := `
              ┏┓     ┏┓                  
              ┃┃    ┏┛┗┓                 
    ┏━━┓┏━━┓┏━┛┃┏━━┓┗┓┏┛┏┓ ┏┓┏━━┓┏━━┓    
    ┃┏━┛┃┏┓┃┃┏┓┃┃┏┓┃ ┃┃ ┃┃ ┃┃┃┏┓┃┃┏┓┃    
    ┃┗━┓┃┗┛┃┃┗┛┃┃┃━┫ ┃┗┓┃┗━┛┃┃┗┛┃┃┃━┫    
    ┗━━┛┗━━┛┗━━┛┗━━┛ ┗━┛┗━┓┏┛┃┏━┛┗━━┛    
                        ┏━┛┃ ┃┃          
                        ┗━━┛ ┗┛          
    `

    title := lipgloss.NewStyle().
        Foreground( m.config.Colors.Foreground ).
        Align(lipgloss.Center).
        Render(titleText)

    var menuOptions strings.Builder

    maxOptionWidth := 0
    for _, option := range m.menuOptions {
        if len(option) > maxOptionWidth {
            maxOptionWidth = len(option)
        }
    }
    maxOptionWidth += 4 


    for i, option := range m.menuOptions {
        optionStyle := lipgloss.NewStyle().
            Padding(0, 1).
            Width(maxOptionWidth).
            Align(lipgloss.Center)

        if m.cursor == i {
            optionStyle = optionStyle.
                Background( m.config.Colors.Color1 ).
                Foreground( m.config.Colors.Foreground )
        } else {
            optionStyle = optionStyle.
                Foreground( m.config.Colors.Foreground )
        }

        centeredOption := lipgloss.NewStyle().
            Width(m.width).
            Align(lipgloss.Center).
            Render(optionStyle.Render(option))

        menuOptions.WriteString(centeredOption + "\n")
    }

    helpText := lipgloss.NewStyle().
        Foreground( m.config.Colors.Color7 ).
        Width(m.width).
        Align(lipgloss.Center).
        MarginTop(1).
		Render("↑/k: up • ↓/j: down • enter: select • q: quit")

    content := lipgloss.JoinVertical(
        lipgloss.Center,
        title,
        "\n",
        menuOptions.String(),
        helpText,
    )

    return menuWin.Render(content)
}


func ( m MenuModel ) handleResize( height, width int ) MenuModel {
	
	m.height = height
	m.width = width
	return m
}
