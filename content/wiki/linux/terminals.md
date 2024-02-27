# Terminal Emulators

A terminal emulator is a program used to interact directly with your OS via text input. Named as "emulators" because they are a software version of the old fashioned [hardware terminals](https://en.wikipedia.org/wiki/Computer_terminal). Most OSes will come with a terminal emulator, Windows includes [Windows Terminal](https://github.com/microsoft/terminal) and [Powershell](https://en.wikipedia.org/wiki/PowerShell), macOS includes [Terminal](https://en.wikipedia.org/wiki/Terminal_(macOS)). Linux desktop environments generally have their own terminal emulators, for example [Xfce Terminal](https://docs.xfce.org/apps/terminal/start), [Gnome Terminal](https://en.wikipedia.org/wiki/GNOME_Terminal) or [Konsole](https://en.wikipedia.org/wiki/Konsole) (for KDE).

This page mainly discusses Linux terminal emulators. Terminal emulators from one distro can be installed on another, and there are a variety of "third party" options to choose from as well.

## Alacritty

[Alacritty](https://alacritty.org/) is a terminal emulator written in Rust, designed for speed and easy customisability. 

### Configuration

To change the default configuration, edit the config file (`~/.config/alacritty/alacritty.yml`). If this config file doesn't exist, the default one (normally located at `/usr/share/doc/alacritty/example/alacritty.yml`) can be copied to the config folder, or one can be created.

Note: Alacritty recently changed from using `yml` to `toml` for config. I will include both types here incase an older version is being used.

#### Useful Config Tips

To change cursor from block to beam:

**yml**

``` yml
cursor:
  style:
    shape: Beam
```

**toml**

``` toml
[cursor.style]
shape = "Beam"
```

To change default window size on opening:

**yml**

``` yml
window:
  dimensions:
    columns: 100
    lines: 25
```

**toml**

``` toml
[window.dimensions]
columns = 100
lines = 25
```

By default, links are clickable without holding down `ctrl`. This can be annoying, to enable clicking links only when holding control, use:

**yml**

``` yml
hints:
  enabled:
  - regex: "(mailto:|gemini:|gopher:|https:|http:|news:|file:|git:|ssh:|ftp:)\
            [^\u0000-\u001F\u007F-\u009F<>\"\\s{-}\\^⟨⟩`]+"
    command: xdg-open
    post_processing: true
    mouse:
      enabled: true
      mods: Control
```

**toml**

``` toml
[[hints.enabled]]
command = "xdg-open"
post_processing = true
regex = "(mailto:|gemini:|gopher:|https:|http:|news:|file:|git:|ssh:|ftp:)[^\u0000-\u001F\u007F-<>\"\\s{-}\\^⟨⟩`]+"

[hints.enabled.mouse]
enabled = true
mods = "Control"
```

Source for this config trick: https://mitjafelicijan.com/alacritty-open-links-with-modifier.html

### Usage

Vi mode

### Troubleshooting

#### Wayland Titlebar

In order to get a titlebar on Wayland, it is neccessary to set the `WAYLAND_DISPLAY` environmental variable to empty. This can be done in the same command that launches Alacritty: 

```
env -u WAYLAND_DISPLAY alacritty
```

#### Ueberzug Issues

Ueberzug is a program to display images "in" any terminal emulator, It works by calculating the position a window should be in, then opening the image in a seperate borderless window and "locking" it to the terminal. Ueberzug is used in a variety of programs, for example [ranger file manager](fileAccess.md#ranger), and my own [GoTube YouTube terminal client](/Wiki/MyPrograms/gotube.md).  
By default, Alacritty will "pad" the window only on the right and the bottom. This padding occurs when the window size is not an exact multiple of the cell size. Ueberzug does not detect this, and thus the images will be misaligned. Enabling "Dynamic Padding" solves this issue by making alaceritty pad evenly on all sides. 

To enable dynamic padding, add the following to `alacritty.toml`

**yml**

``` yml
window:
  dynamic_padding: true
```

**toml**

``` toml
[window]
dynamic_padding = true
```

## Xfce Terminal

[Xfce Terminal](https://gitlab.xfce.org/apps/xfce4-terminal) comes by default in the [Xfce4 Desktop Environment](xfce4.md).

## Cool Retro Term

[Cool Retro Term](https://github.com/Swordfish90/cool-retro-term) is a terminal emulator designed to look like a CRT display, written using the [QT5 UI library](https://doc.qt.io/qt.html#qt5).
