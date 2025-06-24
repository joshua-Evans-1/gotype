package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"gotype/models"
)

func main() {
	p := tea.NewProgram(
		models.NewAppModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, err := p.Run(); err != nil {
		print( fmt.Errorf( "error : %v", err ) )
	}
}

