// models/messages.go

package models

type ChangeViewMsg struct { 
	View string
}

type StartGameMsg struct {
	err error
}

type StopGameMsg struct {
	err error
}

type EditorFinishedMsg struct { 
	err error
}


