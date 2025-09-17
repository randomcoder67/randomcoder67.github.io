# Useful Snippets

Most of these aren't mine, sourced where possible

## Remove Duplicates While Maintaining Order

`cat -n file_name | sort -uk2 | sort -n | cut -f2-`

### Explanation

1. Use `cat -n` to prepend line numbers
2. Use `sort -u` remove duplicate data (`-k2` says 'start at field 2 for sort key')
3. Use `sort -n` to sort by prepended number
4. Use `cut` to remove the line numbering (`-f2-` says 'select field 2 till end')

[Credit to Digital Trauma on StackOverflow](https://stackoverflow.com/a/20639730)

## Firefox Fake Fullscreen - Useful For Tiling WMs

Firefox will not take up the entire screen when a video enters fullscreen, just the space of the Firefox window

Go to `about:config` and set `full-screen-api.ignore-widgets` to `true`

[Credit to You_Yew_Ewe on reddit](https://www.reddit.com/r/i3wm/comments/k2tszq/is_there_a_way_to_get_firefox_to_do_fullscreen/)

## Stress Test USB Drive

`fio --name=usbseq --filename=test.test --size=1G --rw=readwrite --bs=1M --direct=1 --numjobs=1 --iodepth=16 --group_reporting --time_based --runtime=36000`

## Build Conky with Lua Support

* Also makes it install to user `~/.local` directory

`cmake -S . -B build --fresh -DBUILD_X11=ON -DBUILD_WLAN=ON -DBUILD_LUA_CAIRO=ON -DBUILD_LUA_CAIRO_XLIB=ON -DCMAKE_INSTALL_PREFIX="$HOME/.local"`

## Linux Samba Setup

* This example uses the apt package manager, so will work on Debian (including Raspbian) or Ubuntu

1. `sudo apt install samba samba-common-bin`
2. `mkdir ~/Videos/shared`
3. `sudo nano /etc/samba/smb.conf` - See below for file contents
4. `sudo smbpasswd -a <username>`
5. `sudo systemctl restart smbd`

### /etc/samba/smb.conf

``` conf
[RaspberryPiSamba]
path = /home/<username>/Videos/shared
writeable = yes
browseable = yes
public = no
```

[Credit to Py My Life Up](https://pimylifeup.com/raspberry-pi-samba/)

## Edit /sys/.../brightness File Without Sudo Password

`sudo visudo -f /etc/sudoers.d/brightness`
`ALL ALL = (ALL) NOPASSWD: /usr/bin/tee /sys/class/backlight/intel_backlight/brightness`

Edit with:  
`echo 200 | sudo /usr/bin/tee "/sys/class/backlight/intel_backlight/brightness"`

[Linux Mint Forum Post](https://forums.linuxmint.com/viewtopic.php?t=274060)

## Access MediaWiki API Sandbox

Append:  
`wiki/Special:ApiSandbox`  
To the end of the website URL

## Search the Whole of GitHub

Append:  
/search?q=search_terms

