# Contributing to gotype

First off, thank you for considering contributing to **gotype**!
Contributions of all kinds are welcome — from fixing bugs and improving documentation to adding features or suggesting improvements.

This document outlines the basic guidelines for contributing.

---

## Ways to Contribute

* Reporting bugs: Open an issue if you find unexpected behavior.
* Suggesting features: Got an idea? File a feature request issue.
* Improving documentation: Even small fixes help a lot!
* Code contributions: Submit a pull request (PR) with fixes or new features.

---

## Development Setup

1. Fork and clone the repository
2. Build the project
3. Run gotype

---

## Coding Guidelines

* Follow idiomatic Go practices (use `go fmt` and `go vet`).
* Keep code modular and easy to test.
* Write clear commit messages:

  * `fix: correct typo in main.go`
  * `feat: add new theme option`
  * `docs: update installation steps`

loosely follow the [Conventional Commits](https://www.conventionalcommits.org/) format.

id prefer if youd use the same coding styling as me but not required
```go codeStyleExample.go
func ChangeView( view string ) tea.Cmd { 
	return func() tea.Msg { 
		return ChangeViewMsg{ View: view }
	}
}

```

* tabstops 4 
* no new line braces
* spaces between braces
* snake case

---

## Pull Request Process

1. Create a new branch for your change:

   ```
   git checkout -b feat/new-feature
   ```

2. Make your changes and commit with a descriptive message.

3. Push your branch and open a PR against the `main` branch.

4. Ensure your PR passes build checks (if any).

5. Wait for review and feedback!

---

## Communication

* Issues and pull requests are the main way to communicate.
* Be respectful and constructive — we’re all here to learn and improve.

---

## License

By contributing to **gotype**, you agree that your contributions will be licensed under the same license as the project (see [LICENSE](./LICENSE)).
