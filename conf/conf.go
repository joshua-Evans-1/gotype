// conf/conf.go

package conf

import (
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

type Config struct {
	Colors 			Colors
	TimerTimeout 	float64
	Mode 			string
	Tabstops		int
	NewlineBraces	bool
	Case 			string
}

func NewConfig() ( Config ) {

	c := Config{}
	c.Colors = c.Colors.Defaults()

	userConfigDir, confErr := os.UserConfigDir()

	if confErr == nil {
		return c		
	}

	configPath := filepath.Join( userConfigDir, "gotype" )

	viper.SetConfigName( "gotype" )
	viper.SetConfigType( "ini" )
	viper.AddConfigPath( configPath )
	if err := viper.ReadInConfig(); err != nil {
		log.Printf( "No config file found, using default: %v", err )
	}
	
	c = c.readColors( *viper.GetViper() )

	c.TimerTimeout = viper.GetFloat64( "GAME.timeout" )
	c.Mode = viper.GetString( "GAME.mode" )
	c.Tabstops = viper.GetInt( "GAME.tabstops" )
	c.NewlineBraces = viper.GetBool( "GAME.newline_braces" )
	c.Case = viper.GetString( "GAME.case" )

	return c
}

func ( c Config ) readColors( viper viper.Viper ) Config {
	
	c.Colors.Foreground = lipgloss.Color( viper.GetString( "COLORS.foreground" ) ) 
	c.Colors.Background = lipgloss.Color( viper.GetString( "COLORS.background" ) ) 
	c.Colors.Cursor     = lipgloss.Color( viper.GetString( "COLORS.cursor" ) ) 

	c.Colors.Color0  = lipgloss.Color( viper.GetString( "COLORS.color0" ) ) 
	c.Colors.Color1  = lipgloss.Color( viper.GetString( "COLORS.color1" ) ) 
	c.Colors.Color2  = lipgloss.Color( viper.GetString( "COLORS.color2" ) ) 
	c.Colors.Color3  = lipgloss.Color( viper.GetString( "COLORS.color3" ) ) 
	c.Colors.Color4  = lipgloss.Color( viper.GetString( "COLORS.color4" ) ) 
	c.Colors.Color5  = lipgloss.Color( viper.GetString( "COLORS.color5" ) ) 
	c.Colors.Color6  = lipgloss.Color( viper.GetString( "COLORS.color6" ) ) 
	c.Colors.Color7  = lipgloss.Color( viper.GetString( "COLORS.color7" ) ) 
	c.Colors.Color8  = lipgloss.Color( viper.GetString( "COLORS.color8" ) ) 
	c.Colors.Color9  = lipgloss.Color( viper.GetString( "COLORS.color9" ) ) 
	c.Colors.Color10 = lipgloss.Color( viper.GetString( "COLORS.color10" ) ) 
	c.Colors.Color11 = lipgloss.Color( viper.GetString( "COLORS.color11" ) ) 
	c.Colors.Color12 = lipgloss.Color( viper.GetString( "COLORS.color12" ) ) 
	c.Colors.Color13 = lipgloss.Color( viper.GetString( "COLORS.color13" ) ) 
	c.Colors.Color14 = lipgloss.Color( viper.GetString( "COLORS.color14" ) ) 
	c.Colors.Color15 = lipgloss.Color( viper.GetString( "COLORS.color15" ) ) 
	
	return c

}

