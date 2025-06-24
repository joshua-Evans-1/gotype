// models/MenuModel.go

package models

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuModel struct {
	width		int
	height		int
	menuOptions	[]string
	cursor		int
	selected 	string
}

func ( m MenuModel ) Init() tea.Cmd {
	return nil
}

func ( m MenuModel ) HandleInput( msg string ) ( tea.Model, tea.Cmd ) {
	switch( msg ) {
		case "j", "down":
		case "k", "up":
		case "enter", "l":
		case "q", "esc":
			return m, tea.Quit
	}
	return m, nil
}

func ( m MenuModel ) Update( msg tea.Msg ) ( tea.Model, tea.Cmd ) {
	return m, nil
}


func (m MenuModel) View() string {
    // Full-width container that respects terminal boundaries
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
        Foreground(lipgloss.Color("#7D56F4")).
        Align(lipgloss.Center).
        Render(titleText)

    var menuOptions strings.Builder

    // Calculate maximum option width for centering
    maxOptionWidth := 0
    for _, option := range m.menuOptions {
        if len(option) > maxOptionWidth {
            maxOptionWidth = len(option)
        }
    }
    // Add padding to max width
    maxOptionWidth += 4 

    // Color definitions
    selectedBg := "#7D56F4"
    selectedFg := "#FFFFFF"
    defaultFg := "#DDDDDD"

    // Render menu options with dynamic width
    for i, option := range m.menuOptions {
        optionStyle := lipgloss.NewStyle().
            Padding(0, 1).
            Width(maxOptionWidth).
            Align(lipgloss.Center)

        if m.cursor == i {
            optionStyle = optionStyle.
                Background(lipgloss.Color(selectedBg)).
                Foreground(lipgloss.Color(selectedFg))
        } else {
            optionStyle = optionStyle.
                Foreground(lipgloss.Color(defaultFg))
        }

        // Center each option individually within the available space
        centeredOption := lipgloss.NewStyle().
            Width(m.width).
            Align(lipgloss.Center).
            Render(optionStyle.Render(option))

        menuOptions.WriteString(centeredOption + "\n")
    }

    // Help text with dynamic width
    helpText := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#777777")).
        Width(m.width).
        Align(lipgloss.Center).
        MarginTop(1).
        Render("↑/k: up • ↓/j: down • enter: select • q: quit")

    // Combine all components with proper spacing
    content := lipgloss.JoinVertical(
        lipgloss.Center,
        title,
        "\n",
        menuOptions.String(),
        helpText,
    )

    return menuWin.Render(content)
}


