// /models/commands.go

package models

import ( 
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"os/exec"

)

func ChangeView( view string ) tea.Cmd { 
	return func() tea.Msg { 
		return ChangeViewMsg{ View: view }
	}
}

func StartGame() tea.Cmd {
	return func () tea.Msg {
		return StartGameMsg{}
	}
}

func StopGame() tea.Cmd {
	return func () tea.Msg {
		return StopGameMsg{}
	}
}
func openEditor() tea.Cmd {
	editor := os.Getenv( "EDITOR" )
	home := os.Getenv( "HOME" )
	if editor == "" {
		editor = "vim"
	}
	c := exec.Command( editor, home + "/.config/gotype/gotype.conf" ) 
	return tea.ExecProcess( c, func( err error ) tea.Msg {
		return EditorFinishedMsg{ err }
	})
}



