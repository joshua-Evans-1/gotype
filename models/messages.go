// models/messages.go

package models

type ChangeViewMsg struct { 
	View string
}

type EditorFinishedMsg struct { 
	err error
}


