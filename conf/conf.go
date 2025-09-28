// conf/conf.go

package conf

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
)

type Config struct {
	Colors 	Colors
}

func NewConfig() ( Config ) {

	c := Config{}

	c.Colors.Foreground = lipgloss.Color( "#f09fa3" )
	c.Colors.Background = lipgloss.Color( "#0d090b" )
	c.Colors.Cursor = lipgloss.Color( "#f09fa3" )

	c.Colors.Color0 = lipgloss.Color( "#0d090b" )
	c.Colors.Color1 = lipgloss.Color( "#A81B32" )
	c.Colors.Color2 = lipgloss.Color( "#912633" )
	c.Colors.Color3 = lipgloss.Color( "#B82437" )
	c.Colors.Color4 = lipgloss.Color( "#C42A3D" )
	c.Colors.Color5 = lipgloss.Color( "#D43244" )
	c.Colors.Color6 = lipgloss.Color( "#E23C4C" )
	c.Colors.Color7 = lipgloss.Color( "#f09fa3" )
	c.Colors.Color8 = lipgloss.Color( "#a86f72" )
	c.Colors.Color9 = lipgloss.Color( "#A81B32" )
	c.Colors.Color10 = lipgloss.Color( "#912633" )
	c.Colors.Color11 = lipgloss.Color( "#B82437" )
	c.Colors.Color12 = lipgloss.Color( "#C42A3D" )
	c.Colors.Color13 = lipgloss.Color( "#D43244" )
	c.Colors.Color14 = lipgloss.Color( "#E23C4C" )
	c.Colors.Color15 = lipgloss.Color( "#f09fa3" )

	userConfigDir, confErr := os.UserConfigDir()

	if confErr == nil {
		return c		
	}


	filepath.Join( userConfigDir, "gotype", "gotype.conf" )
	return c
}
