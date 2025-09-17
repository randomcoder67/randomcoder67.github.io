# mpv

\infobox
title: mpv
img: /images/programScreenshots/mpv-example.png
caption: A YouTube video playing in mpv
Website|[mpv.io](https://mpv.io/)
GitHub|[mpv GitHub](https://github.com/mpv-player/mpv/)
Language|C (configured in Lua)
Category|Media Player

mpv is a media player, designed to be configured from the command line instead of via a GUI interface. An mpv window only contains the video, and the "OSC" (on screen controller), there are no menus for loading videos, or changing settings.

mpv can play video and audio, either from local files or from network streams. It includes integration with `yt-dlp`, to allow playback of YouTube videos.

As mpv is highly configurable, it is often used in scripts, for example a simple music player can be made by mapping a hotkey to `mpv ~/Music/somePlaylist`. Unix sockets can be used to communicate with a running mpv instance, allowing control of the playback without focusing the window (you need not even have a window open).

## mpv.conf

`~/.config/mpv/mpv.conf` is used for general configuration. Details about some useful/interesting options are listed below:

| Function                               | Setting                               | Explanation                                                                                                                                                       |
| :------------------------------------: | :-----------------------------------: | :---------------------------------------------------------------------------------------------------------------------------------------------------------------: |
| Enabled hardware decode for Intel GPUs | `hwdec=vaapi`                         | See: [Intel Quick Sync Video Wikipedia](https://en.wikipedia.org/wiki/Intel_Quick_Sync_Video) for details about codec support                                     |
| Set Default Window Size to 30%         | `geometry=30%`                        | By default videos will open at their true size. This makes them smaller so they don't take over the entire screen                                                 |
| Save Video Progress                    | `save-position-on-quit`               | Video will reopen at the point you left off (unless `--no-resume-playback` used)                                                                                  |
| Set Screenshot Directory               | `screenshot-directory=~/Pictures/mpv` | Uses `~/Pictures/mpv`                                                                                                                                             |
| Always Open Window                     | `force-window=yes`                    | Causes a blank window to open when file without video is played (e.g. audio file with no album art)                                                               |
| Fix Issue With Video/Audio Desync      | `hr-seek=yes`                         | Fixes an issue when playing video and audio from different network streams (e.g. most YouTube formats), where rewinding will cause audio desync and video speedup |
| Set Max Cache                          | `demuxer-max-bytes=50MiB`             | Limits the cache to 50 MiB                                                                                                                                        |
| Set Max Rewind Cache                   | `demuxer-max-back-bytes=20MiB`        | Limits the rewind cache to 20 MiB                                                                                                                                 |
| Turn Default OSC Off                   | `osc=no`                              | Removes default OSC as I use osc-tethys                                                                                                                           |
| Enable Border                          | `border=yes`                          | Ensures the window manager borders are used                                                                                                                       |
| Set OSC Font                           | `osc-font=Roboto`                     | Uses Roboto                                                                                                                                                       |
| Set OSC Font Size                      | `osc-font-size=35`                    | Uses font size 35                                                                                                                                                 |
| Turn Auto Window Resize Off            | `auto-window-resize=no`               | Honestly can't remember what this does                                                                                                                            |
| Turn OSD Scale Off (1)                 | `script-opts=osc-vidscale=no`         | Stops (mainly) the scaling of the OSC depending on window size                                                                                                    |
| Turn OSD Scale Off (2)                 | `osd-scale-by-window=no`              |                                                                                                                                                                   |


[AudioFiles]
profile-cond = (filename:match"%.mp3$" or filename:match"%.m4a$") ~= nil
#title='${metadata/title} - ${metadata/artist}'
geometry=25%
#<tempo|pitch|both|none>

[extension.vpy]
demuxer-lavf-format=vapoursynth


## input.conf

`~/.config/mpv/input.conf` is used to set keybinds for mpv. My keybinds are listed below (not all of these are in `input.conf`, as some are the default):

### Video Navigation

| Function          | Shortcut                       | Explanation                               |
| :---------------: | :----------------------------: | :---------------------------------------: |
| Pause/Play        | `Space`                        |                                           |
| Skip 5 Seconds    | `Right`                        |                                           |
| Rewind 5 Seconds  | `Left`                         | Use negative number to rewind (`seek -5`) |
| Skip 30 Seconds   | `Up`                           |                                           |
| Rewind 30 Seconds | `Down`                         | `seek -30` is used to rewind              |
| Next Chapter      | `Shift + Right`                |                                           |
| Previous Chapter  | `Shift + Left`                 |                                           |
| Next Video        | `Enter` or `Shift + n`         |                                           |
| Previous Video    | `Shift + Enter` or `Shift + p` |                                           |

### Speed

| Function                | Shortcut | Explanation                                              |
| :---------------------: | :------: | :------------------------------------------------------: |
| Increase speed by `0.1` | `s`      |                                                          |
| Decrease speed by `0.1` | `d`      | Use negative number to decrease speed (`add speed -0.1`) |
| Reset Speed             | `r`      |                                                          |
| Toggle Speed 1x and 2x  | `g`      |                                                          |
| Show Current Speed      | `x`      |                                                          |

The speed control keybinds are the same as [Video Speed Controller](https://addons.mozilla.org/en-GB/firefox/addon/videospeed/), a popular addon to control HTML video playback  

### Other

| Function          | Shortcut    | Explanation               |
| :---------------: | :---------: | :-----------------------: |
| Screenshot        | `Shift + s` | Saves to `~/Pictures/mpv` |
| Toggle Mute       | `m`         |                           |
| Show Clock        | `t`         |                           |
| Toggle Clock      | `Shift + t` |                           |
| Show Chapters     | `c`         |                           |
| Show Playlist     | `u`         |                           |
| Show Metadata     | `Shift + m` |                           |
| Show OSC          | `/`         |                           |
| Toggle OSC        | `Shift + /` |                           |
| Toggle Mono Audio | `Shift + v` |                           |
| Rotate Video 90°  | `Shift + r` |                           |
| Set Geometry 60%  | `Shift + f` | "Half Fullscreen"         |
