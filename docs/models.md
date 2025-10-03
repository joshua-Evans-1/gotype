
# models

## AppModel.go

Defines the top-level application model. Coordinates state between sub-models (menu, game, status bar) and manages global messages/commands.

## GameModel.go

Handles the typing game logic. Tracks test progress, words, WPM, accuracy, and timing. Responds to user keystrokes and updates game state accordingly.

## MenuModel.go

Implements the main menu UI. Provides options to start a test, configure settings, or quit. Routes user selections to the appropriate sub-models.

## StatusBarModel.go

Defines the status bar displayed at the bottom of the TUI. Shows contextual information such as time remaining, WPM, and accuracy.

## commands.go

Contains Bubble Tea commands for asynchronous actions. Examples include starting/stopping the timer, loading word lists, or switching themes.

## messages.go

Defines message types used in the Bubble Tea update loop. Messages are dispatched between models to handle events like keystrokes, timer ticks, or config changes.

## timer.go

Implements a game timer. Sends periodic tick messages to update the game model and stop the test when time runs out.

