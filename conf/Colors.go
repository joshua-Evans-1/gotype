// conf/Colors.go

package conf

import (
	"github.com/charmbracelet/lipgloss"
)

type Colors struct { 

	Foreground	lipgloss.Color
	Background	lipgloss.Color
	Cursor 		lipgloss.Color

	Color0		lipgloss.Color
	Color1		lipgloss.Color
	Color2		lipgloss.Color
	Color3		lipgloss.Color
	Color4		lipgloss.Color
	Color5		lipgloss.Color
	Color6		lipgloss.Color
	Color7		lipgloss.Color
	Color8		lipgloss.Color
	Color9		lipgloss.Color
	Color10		lipgloss.Color
	Color11		lipgloss.Color
	Color12		lipgloss.Color
	Color13		lipgloss.Color
	Color14		lipgloss.Color
	Color15		lipgloss.Color
}

func ( c Colors ) Defaults(  ) Colors {
	
	c.Foreground = lipgloss.Color( "#f09fa3" )
	c.Background = lipgloss.Color( "#0d090b" )
	c.Cursor = lipgloss.Color( "#f09fa3" )

	c.Color0 = lipgloss.Color( "#0d090b" )
	c.Color1 = lipgloss.Color( "#A81B32" )
	c.Color2 = lipgloss.Color( "#912633" )
	c.Color3 = lipgloss.Color( "#B82437" )
	c.Color4 = lipgloss.Color( "#C42A3D" )
	c.Color5 = lipgloss.Color( "#D43244" )
	c.Color6 = lipgloss.Color( "#E23C4C" )
	c.Color7 = lipgloss.Color( "#f09fa3" )
	c.Color8 = lipgloss.Color( "#a86f72" )
	c.Color9 = lipgloss.Color( "#A81B32" )
	c.Color10 = lipgloss.Color( "#912633" )
	c.Color11 = lipgloss.Color( "#B82437" )
	c.Color12 = lipgloss.Color( "#C42A3D" )
	c.Color13 = lipgloss.Color( "#D43244" )
	c.Color14 = lipgloss.Color( "#E23C4C" )
	c.Color15 = lipgloss.Color( "#f09fa3" )
	return c
}
