# Other Configs

## GTK 3

`~/.config/gtk-3.0/gtk.css`

``` css
#xfwm-tabwin {
-XfwmTabwinWidget-icon-size: 128px;
-XfwmTabwinWidget-preview-size: 256px; }
```

Sets the default icon and window preview sizes for the `Alt + Tab` window switcher.

## npm

`~/.config/npm/npmrc`

```
prefix=${XDG_DATA_HOME}/npm
cache=${XDG_CACHE_HOME}/npm
init-module=${XDG_CONFIG_HOME}/npm/config/npm-init.js
```

Makes npm use XDG compliant directories instead of putting stuff in the home folder

## tmux

`~/.config/tmux/tmux.conf`

| Function                  | Shortcut             | Setting                              | Explanation                                                                                                                   |
| :-----------------------: | :------------------: | :----------------------------------: | :---------------------------------------------------------------------------------------------------------------------------: |
| Switch to Next Pane       | `Shift + Tab`        | `bind-key -n BTab select-pane -t +1` | Panes are the "splits"                                                                                                        |
| Switch to Next Window     | `Ctrl + Tab`         | `bind -n C-Tab select-window -n`     | Same keyboard shortcut as tabs in text editor or browser                                                                      |
| Switch to Previous Window | `Ctrl + Shift + Tab` | `bind -n C-S-Tab select-window -p`   |                                                                                                                               |
| Enable Mouse              | N/A                  | `set -g mouse on`                    | Mouse can be used to drag-resize panes                                                                                        |
| Set `Esc` Delay to 0      | N/A                  | `set -g escape-time 0`               | By default there is a delay after pressing `Esc` to ensure an `Alt` sequence wasn't pressed. I set this to 0 as it's annoying |


## btop

`~/.config/btop/btop.conf`

Notable Settings:

| Function                        | Setting                     | Explanation                                                                                       |
| :-----------------------------: | :-------------------------: | :-----------------------------------------------------------------------------------------------: |
| Set Refresh to 1000ms (1 sec)   | `update_ms = 1000`          |                                                                                                   |
| Set Process Sorting to CPU Lazy | `proc_sorting = "cpu lazy"` | Easier to read and more useful than "cpu direct"                                                  |
| Set Process Per Core to On      | `proc_per_core = True`      | Makes it easier to see processes using a lot of CPU, as using 1 full core will appear as "100.0%" |

## Cool Retro Term

`~/.config/cool-retro-term/cool-retro-term.conf`

``` toml
[QQControlsFileDialog]
favoriteFolders=@Invalid()
height=0
sidebarSplit=132.525
sidebarVisible=true
sidebarWidth=80
width=0
```

## Rofi



## Micro


* rofi
* micro
