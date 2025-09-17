# File Access

File access is one of the most important parts of using a computer. As one of the things you do the most, it is important to have an easy and efficient way to interact with the file system. There are three main types of file management systems, graphical file managers, terminal file managers and command line file managing. This article goes over the basics of each, and some examples of programs/tools. 

## Graphical File Managers

The most common way of managing files, GUI file managers come as standard on most OSes. Windows has Windows Explorer, macOS has Finder, GNOME has Nautilus, Xfce has Thunar etc. The advantage of a GUI file manager is the ability to use the mouse to click and drag files.

### Thunar

[Thunar](https://docs.xfce.org/xfce/thunar/start) is the default file manager included in the [Xfce desktop environment](xfce.md)

## Terminal File Managers

Terminal file managers (also known as TUI file managers) are somewhat of a middle ground between a full GUI file manager, and simply using the command line. They are generally operating using the keyboard only, and have a number of shortcuts to make common tasks easier. Some popular options for terminal file managers are listed below. (not much detail here because I haven't used any of these extensively)

### Ranger

[Ranger](https://github.com/ranger/ranger) is a terminal file manager written in [Python 3](python3.md) 

### nnn

[nnn](https://github.com/jarun/nnn) is a terminal file manager mainly written in [c](c.md)

### lf

[lf](https://github.com/gokcehan/lf) is a terminal file manager written in [Golang](golang.md)

## Command Line File Managing

Command line (CLI) file managing is the most basic way of managing files on a computer. The main program required is simply a [terminal emulator](terminals.md), and the [GNU Core Utils](coreutils.md), such as `cd`, `mv` and `rm`. There are of course a variety of programs which significantly increase the efficiency of command line file management. Some of these are listed here. 

### Fasd

[Fasd](https://github.com/clvv/fasd) is an autojump program, designed to allow quick opening of files and changing directories in the terminal. Fasd uses a metric called "frecency" (frequency and recency) to order the files you access. 

#### Setup

`pacman -S fasd`  
Add `eval "$(fasd --init auto)"` to the bottom of your `~/.bashrc` file. 

#### Basic Usage

Fasd comes with some basic [aliases](bash.md). These are listed below:

`a name` - List all files/directories matching "name", in order of worst to best match  
`s name` - Same as a, but allows the user to select  
`d name` - Same as a, but only for directories  
`f name` - Same as a, but only for files  
`sd name` - Same as d, but allows the user to select  
`sf name` - Same as f, but allows the user to select  
`z name` - Jump to directory that matches "name" most closely  
`zz name` - Same as above, but allow the user to pick from all matches instead of jumping straight to the best match  

#### Advanced Usage (Integrating with Other Programs)

Fasd can be easily incorporated into functions to allow quick opening of files in certain programs. For example, my main text editor is [Mousepad](textEditors.md), so I created this bash function to allow me to quickly open commonly used files in Mousepad:

```
# Open file with Mousepad (fasd)
m () {
    result=$(fasd -fi $@)
    [ "$result" == "" ] && return
    mousepad "$result" & disown
}
```

The functions first used Fasd to list all files matching the given strings, allowing me to select one. It then passes this file to Mousepad and disowns the process, disconnecting it from the terminal.  
It also checks the result is not blank, so as not to open an empty document if no file match was found. 

Fasd can also be used in-situ to search for files. For example, to search for the best directory match to the string "doc", and move "file.txt" there:
 
```
mv file.txt `d doc`
```
 
### trash-cli

[trash-cli](https://github.com/andreafrancia/trash-cli) is a program to implement recycle bin functionality to the terminal, written in Python 3. By default, the GNU Core Util `rm` does not send files to a recycle bin. trash-cli fixes this. By aliasing `rm` to `trash-cli -i`, you will be prompted before deleting a file/folder, and it will only be send to the trash folder `~/.local/share/Trash`. 

#### Basic Usage

`trash-put [-i] file` - Put the given file in the trash, asking for confirmation if `-i` given  
`trash-restore [file]` - List files trashed from current directory (and matching "file" if specified) and allow user to pick files/dirs to restore  
`trash-empty` - Empty the trash folder, **deleting items permanently** and freeing up disk space  
`trash-list` - List all items currently in the trash folder  

### TMSU

[TMSU](https://tmsu.org/) is a file tagging program for the Linux terminal. 

### bat

[bat](https://github.com/sharkdp/bat) is a cat alternative with far more features, including git integration, syntax highlighting and showing non-printable characters. 

### perl-rename

[perl-rename](https://learnbyexample.github.io/learn_perl_oneliners/perl-rename-command.html) is a bulk renaming program using [sed](sed.md) syntax. 
