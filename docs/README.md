# Documentation

## Project structure


<pre style="line-height: 0.5;">

  gotype
  ├── <a href="../LICENSE">LICENSE</a>
  ├── <a href="../README.md">README.md</a>
  ├── <a href="conf.md">conf</a>
  │   ├── <a href="conf.md#Colors.go">Colors.go</a>
  │   └── <a href="conf.md#conf.go">conf.go</a>
  ├── <a href="database.md">database</a>
  │   ├── <a href="database.md#database.go">database.go</a>
  ├── <a href="../go.mod">go.mod</a>
  ├── <a href="../go.sum">go.sum</a>
  ├── <a href="../main.go">main.go</a>
  └── <a href="models.md">models</a>
      ├── <a href="models.md#AppModel.go">AppModel.go</a>
      ├── <a href="models.md#GameModel.go">GameModel.go</a>
      ├── <a href="models.md#MenuModel.go">MenuModel.go</a>
      ├── <a href="models.md#StatusBarModel.go">StatusBarModel.go</a>
      ├── <a href="models.md#commands.go">commands.go</a>
      ├── <a href="models.md#messages.go">messages.go</a>
      └── <a href="models.md#timer.go">timer.go</a>

</pre>


## Frameworks

### Bubble Tea  
Bubble Tea uses the [Elm architecture](https://guide.elm-lang.org/architecture/).  

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case tea.KeyMsg:
    }
    return m, nil
}
```

```go
// Example: simple View function
func (m model) View() string {
    return "Hello, Bubble Tea!"
}
```

---

### Lipgloss  
Lipgloss is used to style components in views.  

```go
style := lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#ff79c6"))

styledText := style.Render("Hello, Lipgloss!")
```

