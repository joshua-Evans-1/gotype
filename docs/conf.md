
# conf

the conf directory holds files for the configuration module


## Colors.go

colors defines a struct and initalization methods for the colorscemes for the stying of the game 
it uses a 16 color palatte and color options for the background, foregorund and cursor

## conf.go

conf defines a struct for the games configuration it reads a file in XDG_CONFIG_HOME/.config/gotype
and initializes the values from that file as well as provides default values for unitialized values

