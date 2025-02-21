# Terminal Directory Navigation

## CLI

In order to speed up moving around directories in the terminal, I have a few shortcuts

### Permenant Directory Aliases

For commonly accessed directories which never change, I simply use aliases for quick access:

``` bash
alias doc='cd ~/Documents/'
alias dow='cd ~/Downloads/'
alias mus='cd ~/Music/'
alias vid='cd ~/Videos/'
alias pic='cd ~/Pictures/'
alias pro='cd ~/Programs/'
alias wor='cd ~/Work/'
alias loc='cd ~/.local/share/'
alias bin='cd ~/.local/bin/'
alias con='cd ~/.config/'
alias bac='cd ~/Downloads/BackupMount/'
```

### directory_bookmakrs

I wrote [my own directory bookmarking script](https://github.com/randomcoder67/Dotfiles-Scripts-and-WM-Setups/blob/main/terminal_scripts/directory_bookmarks.c), based off [apparix](https://github.com/micans/apparix).  
This script is used for non-permenant bookmarks, i.e. directories that I need quick access to at the moment, but aren't a permenant feature of my setup. For example, the `.minecraft` directory for my current MultiMC Minecraft instance, or various directories in `~/.config`

I use aliases to the script in my `.bashrc`:

``` bash
function to {
	cd "$(directory_bookmarks get "$1")"
}
alias bm='directory_bookmarks add'
alias bmr='directory_bookmarks remove'
alias bml='directory_bookmarks list'
alias bmc='directory_bookmarks current'
```

### Quick Jump

I use [Fasd](https://github.com/whjvenyl/fasd) for quickly jumping to recently used folders that aren't bookmarked some other way. 

### Other

`alias cdp='cd - > /dev/null'` - cd to the previous directory, without printing it's name

## TUI

I use [ranger](https://github.com/ranger/ranger) as my TUI file manager.

### Permenant Shortcuts

In ranger, you can use `m` to mark directories with a single letter, and then backtick to jump to that directory

I use these bookmarks, which are as similar as possible to the quick cd aliases:

``` conf
h:/home/user
d:/home/user/Documents
l:/home/user/Downloads
m:/home/user/Music
v:/home/user/Videos
i:/home/user/Pictures
p:/home/user/Programs
w:/home/user/Work
s:/home/user/.local/share
b:/home/user/.local/bin
c:/home/user/.config
e:/home/user/Downloads/BackupMount
```

`h` is a replacement for just typing `cd`, taking you to the home directory

### directory_bookmarks

When using directory_bookmarks.c in the terminal, the bookmarks are also added to ranger, and can be used by typing `go` followed by the bookmark name.

