package main

import (
	"fmt"
	"gotype/database"
	"gotype/models"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	database.ConnDB()
	p := tea.NewProgram(
		models.NewAppModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, err := p.Run(); err != nil {
		print( fmt.Errorf( "error : %v", err ) )
	} 
	sqldb, err := database.DB.DB()
	if err != nil {
		panic( err )
	}
	defer sqldb.Close()
}

