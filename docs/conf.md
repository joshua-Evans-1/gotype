# conf

The `conf` directory contains the **configuration module** for `gotype`.  
It defines how game settings (such as themes, colors, and user preferences) are stored and initialized.

---

## Colors.go

This file defines:

- A `Colors` struct for holding the game’s color scheme  
- Initialization methods for setting up themes  
- Support for a **16-color palette**  
- Options for background, foreground, and cursor colors  

Colors are used throughout the TUI to provide consistent styling.

---

## conf.go

This file defines:

- A `Conf` struct for the game’s configuration state  
- Logic to read user configuration files from:  

$XDG_CONFIG_HOME/.config/gotype

- Default values for uninitialized or missing settings  
- Integration with the `Colors` module to provide a complete configuration

This is the entry point for setting up the game’s config at startup.

