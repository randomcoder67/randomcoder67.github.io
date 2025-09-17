# Alacritty

## Font

I use Roboto Mono as a monospaced font. 
Relevant configuration in `alacritty.toml`:

``` toml
[env]
WINIT_X11_SCALE_FACTOR = "2"

[font]
size = 6

[font.bold]
family = "Roboto Mono"
style = "Bold"

[font.bold_italic]
family = "Roboto Mono"
style = "Bold Italic"

[font.italic]
family = "Roboto Mono"
style = "Italic"

[font.normal]
family = "Roboto Mono"
style = "Regular"
```

The `WINIT_X11_SCALE_FACTOR = "2"` line is used because my laptop has a high DPI display. If you have a 1080p screen or lower, you will likely want to change this to `"1"`, and adjust the font size accordingly.

## Colour Scheme/Theming

I use a slightly modified version of Dracula. Only line of significant note below is the cursor style, whichis set to a "Beam" instead of the default block.

``` toml
[colors.bright]
black = "#7a71aa"
blue = "#bd93f9"
cyan = "#8be9fd"
green = "#50fa7b"
magenta = "#ff79d2"
red = "#ff5555"
white = "#f8f8f2"
yellow = "#f1fa8c"

[colors.cursor]
cursor = "CellForeground"
text = "CellBackground"

[colors.footer_bar]
background = "#282a36"
foreground = "#f8f8f1"

[colors.hints.end]
background = "#282a36"
foreground = "#f1fa8c"

[colors.hints.start]
background = "#f1fa8c"
foreground = "#22212d"

[colors.line_indicator]
background = "None"
foreground = "None"

[colors.normal]
black = "#22212d"
blue = "#9a7cff"
cyan = "#38ffea"
green = "#52ff6d"
magenta = "#ff76c1"
red = "#e64747"
white = "#f8f8f1"
yellow = "#ffff80"

[colors.primary]
background = "#22212d"
bright_foreground = "#ffffff"
foreground = "#f8f8f1"

[colors.search.focused_match]
background = "#ffb86c"
foreground = "#44475a"

[colors.search.matches]
background = "#50fa7b"
foreground = "#44475a"

[colors.selection]
background = "#44475a"
text = "CellForeground"

[colors.vi_mode_cursor]
cursor = "CellForeground"
text = "CellBackground"

[cursor.style]
shape = "Beam"
```

## Other

``` toml
[window]
dynamic_padding = true
```

Enabling dynamic padding affects how the window is rendered when it it's size is not an exact multiple of the "cell" size. With dynamic padding on, padding is added (relatively) evenly on all sides, instead of only at the bottom and right.  
This option also enabled ueberzug to align images properly.

``` toml
[window.dimensions]
columns = 100
lines = 25
```

Sets the default window size. Keep in mind a "cell" is not square.

``` toml
[[hints.enabled]]
command = "xdg-open"
post_processing = true
regex = "(mailto:|gemini:|gopher:|https:|http:|news:|file:|git:|ssh:|ftp:)[^\u0000-\u001F\u007F-<>\"\\s{-}\\^⟨⟩`]+"

[hints.enabled.mouse]
enabled = true
mods = "Control"
```

Enabled clicking on links only when the `Ctrl` key is held down.

``` toml
[[keyboard.bindings]]
chars = "\u001B[27;5;9~"
key = "Tab"
mods = "Control"

[[keyboard.bindings]]
chars = "\u001B[27;6;9~"
key = "Tab"
mods = "Control|Shift"
```

Enabled `Ctrl + Tab` and `Ctrl + Shift + Tab` keyboard shortcuts, which normally don't work in terminals. In my setup this is used for switching tmux windows.

``` toml
[[keyboard.bindings]]
chars = "\u0019"
key = "z"
mods = "Control|Shift"
```

Enabled `Ctrl + Shift + Z` keyboard shortcut, used in Micro for `redo` functionality.
