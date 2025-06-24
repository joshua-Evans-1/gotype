// conf/conf.go

package conf

import (
	"os"

	"github.com/charmbracelet/lipgloss"
)

type Config struct {
	
	BG				lipgloss.Color
	FG				lipgloss.Color
	PrimaryColor	lipgloss.Color
	SecondaryColor	lipgloss.Color
}

func ( c Config ) new_Config() ( Config ) {
	homeDir := os.Getenv("HOME")
	confDir := homeDir + "/.config/gotype"

	os.Mkdir( confDir, 0755 )

	os.Create( confDir + "codetype.conf" )

	return c
}
