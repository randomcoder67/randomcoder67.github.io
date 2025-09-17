# Xfce/Xfwm Keyboard Shortcuts

I try and limit keyboard shortcuts to either programs/commands I want to use very regularly, or tasks that need to be performed without shifting focus from the current window.

## Music

| Function                                      | Shortcut            | Explanation                                                                           |
| :-------------------------------------------: | :-----------------: | :-----------------------------------------------------------------------------------: |
| Play Music from Current Playlist              | `Super + L`         | Current Playlist refers to `~/Music/CurrentPlaylist`                                  |
| Play Music from Current Playlist (Background) | `Super + Shift + L` | "Background" meaning no mpv window                                                    |
| Select Playlist to Play                       | `Super + O`         | Playlists are simply folders in `~/Music`                                             |
| Select Playlist to Play (Background)          | `Super + Shift + O` | The playlist choosing menu opens but will play in background                          |
| Increase Volume                               | `Super + ]`         | Increases by 3% at a time                                                             |
| Decrease Volume                               | `Super + [`         | Decreases by 3% at a time                                                             |
| Toggle Mute                                   | `Super + Shift + ;` |                                                                                       |
| Pause/Play Music                              | `Super + Shift + [` |                                                                                       |
| Next Track in Playlist                        | `Super + =`         |                                                                                       |
| Previous Track in Playlist                    | `Super + -`         |                                                                                       |
| Stop Music                                    | `Super + Shift + -` | Closes the mpv instance                                                               |
| Identify Playing Music                        | `Super + R`         | Uses [songrec](https://github.com/marin-m/SongRec), listens to any audio being played |

## Program Openers

| Function                 | Shortcut                 | Explanation                                                                                                                                 |
| :----------------------: | :----------------------: | :-----------------------------------------------------------------------------------------------------------------------------------------: |
| Default Terminal         | `Super + Return`         | In my config, Alacritty is the default terminal                                                                                             |
| Backup Terminal          | `Super + Shift + Return` | Xfce4 Terminal is the backup terminal, here incase something does wrong with Alacritty                                                      |
| Browser                  | `Super + W`              | Firefox by default in my config                                                                                                             |
| GoTube                   | `Super + S`              | My own YouTube terminal client, see [GoTube] for more info                                                                                  |
| TUI File Manager         | `Super + E`              | ranger by default                                                                                                                           |
| GUI File Manager         | `Super + Shift + E`      | Thunar by default                                                                                                                           |
| Text Editor              | `Super + M`              | Mousepad                                                                                                                                    |
| Text Editor (New Window) | `Super + Shift + M`      |                                                                                                                                             |
| Calculator               | `Super + C`              | Qalculate terminal version                                                                                                                  |
| htop                     | `Super + H`              | Resource Monitor                                                                                                                            |
| btop                     | `Super + B`              | Pretty Resource Monitor                                                                                                                     |
| NetHogs                  | `Super + N`              | Network traffic Monitor                                                                                                                     |
| pulsemixer               | `Super + G`              | Audio Control (this is set to `Super + G` because that brings up the "Game Bar" in Windows, which has volume control, and I got used to it) |
| cava visualiser          | `Super + U`              |                                                                                                                                             |
| System Menu              | `Super + X`              |                                                                                                                                             |

## Utilities

| Function           | Shortcut     | Explanation                                                                                                                   |
| :----------------: | :----------: | :---------------------------------------------------------------------------------------------------------------------------: |
| Rofi Launcher      | `Super`      | A launcher program using Rofi, allows opening programs, shortcuts or files (see rofiLauncher)                                 |
| Search Files       | `Super + F`  | Opens a Rofi script which searches all files (in `/home/user`)                                                                |
| Bookmarks          | `Super + Q`  | Bookmark manager, use `Enter` to copy to clipboard, `Shift + Enter` to open (if URL), `Ctrl + A` to add, `Ctrl + W` to remove |
| Emoji Picker       | `Super + I`  |                                                                                                                               |
| Colour Picker      | `Super + Y`  | Uses XColor, displays the colour hex as a notification and copies to clipboard (without the '#')                              |
| Screenshot Utility | `Super + \`` | Opens Xfce Screenshooter                                                                                                      |
| Help Info          | `Super + V`  |                                                                                                                               |

## System

| Function                                | Shortcut            | Explanation                                                                          |
| :-------------------------------------: | :-----------------: | :----------------------------------------------------------------------------------: |
| Increase Brightness                     | `Super + '`         |                                                                                      |
| Decrease Brightness                     | `Super + #`         |                                                                                      |
| Clear All Notifications                 | `Super + Shift + C` | Brute force, just runs `pkill xfce4-notify`, it restarts when next notification sent |
| Logout                                  | `Ctrl + Alt + Del`  | Opens Xfce logout menu                                                               |
| Expand Left Window in Horizontal Split  | `Super + K`         |                                                                                      |
| Expand Right Window in Horizontal Split | `Super + J`         |                                                                                      |

## Window Manager Functions

| Function                          | Shortcut                 | Explanation                                                                 |
| :-------------------------------: | :----------------------: | :-------------------------------------------------------------------------: |
| Cycle Windows                     | `Alt + Tab`              |                                                                             |
| Cycle Windows of Same Application | `Alt + \``               |                                                                             |
| Show Desktop                      | `Super + D`              | Pressing again will restore previous open windows                           |
| Maximise/Restore Window           | `Super + Shift_R`        | Only works with `Shift_R`, the `Shift` key under Enter                      |
| Minimise Window                   | `Super + Ctrl + Shift_R` | Only works with `Shift_R`, the `Shift` key under Enter                      |
| Close Window/Logout               | `Super + Shift + Q`      | Will open logout dialog if no window focused                                |
| Resize Window                     | `Super + Shift + R`      | Move the mouse to resize from bottom right corner                           |
| Show Window on all Workspaces     | `Super + Shift + X`      |                                                                             |
| Keep Window on Top                | `Super + Shift + F`      |                                                                             |
| Next Workspace                    | `Super + +`              |                                                                             |
| Previous Workspace                | `Super + -`              |                                                                             |
| Move Window to Next Workspace     | `Super + Ctrl + .`       |                                                                             |
| Move Window to Previous Workspace | `Super + Ctrl + ,`       |                                                                             |
| Move Window to Upper Monitor      | `Super + Ctrl + #`       | I have these keys mapped to "Upper" and "Lower" as that is my monitor setup |
| Move Window to Lower Monitor      | `Super + Ctrl + '`       |                                                                             |
| Add New Workspace                 | `Super + Alt + .`        |                                                                             |
| Remove Last Workspace             | `Super + Alt + ,`        | Removes the last workspace numerically                                      |
| Snap Window to Left               | `Super + Left`           |                                                                             |
| Snap Window to Right              | `Super + Right`          |                                                                             |
| Snap Window to Top                | `Super + Up`             |                                                                             |
| Snap Window to Bottom             | `Super + Down`           |                                                                             |
| Snap Window to Bottom Left        | `Super + Fn + Left`      | Mapped to `End`                                                             |
| Snap Window to Top Left           | `Super + Fn + Up`        | Mapped to `PageUp`                                                          |
| Snap Window to Bottom Right       | `Super + Fn + Down`      | Mapped to `PageDown`                                                        |
| Snap Window to Top Right          | `Super + Fn + Right`     | Mapped to `Home`                                                            |

`Super+P` Plugging/Unplugging from Monitor  
`Super+I` Link Markdown Files  
`Super+Shift+]` Delete Song if playing New Music  
`Super+Shift+=` Save Song to toDownload.txt and Delete if playing New Music  
