# Linux Distributions

A Linux distribution (distro) is an OS made from the Linux kernel, with a package manager, and often a desktop environment included. 

(Note: This article only covers floating window managers, see [Tiling Window Managers](tilingwm.md) for more detail about the other kind of window manager)

## Picking a Distro

It is common to spend a long time mulling over which distro to choose, or constantly change between distros (known as "distro-hopping"). A lot of the time, differences between distros are minimal, and features of one can be added to another with relative ease. For this reason, the only two real concerns when picking a distro are the package manager/style of updates, and which Window Manager/Desktop Environment you prefer. 

There are three "main" Linux distributions, Arch Linux, Debian and Fedora. Each of these comes with a different package manager (program used to install other programs), and has a different focus when it comes to which versions of packages are included. This article contains an overview of the three different main distros, as well as some examples of popular distros based on these core three. 

## Arch Linux

[Arch Linux](https://archlinux.org/) is a minimalist, rolling release Linux distro. Rolling release refers to the update schedule for the distro. Arch does not have distinct "versions", but rather is constantly being updated as the packages themselves are updated. This means you (generally) get the most up to date version of every program on your system. This has obvious benefits, but can also lead to a less stable OS than non rolling release distros. 

Arch is somewhat infamous for being difficult to install, and home to ["Linux Elitists"](https://knowyourmeme.com/memes/btw-i-use-arch). The install process for Arch Linux is a lot more difficult than most Distros (or Windows), requiring commands to be entered at a terminal, instead of a GUI install process. There is an [excellent guide on the Arch Wiki](https://wiki.archlinux.org/title/installation_guide) to help first time users install the distro, however it isn't recommended unless you have at least *some* experience in Linux beforehand. 

The "Linux Elitists" idea is less true. As Arch is a "more difficult" distro, users will often expect a higher level of base knowledge from people asking questions, however if you are a true newbie there are plenty of resources available to get help with the distro without any expectation of base knowledge (such as the [Arch subreddit](https://old.reddit.com/r/archlinux/)). 

Arch uses [pacman](https://wiki.archlinux.org/title/pacman) for package management, which contains most of the packages you will need. For the programs that are not available through pacman, the [Arch User Repository](https://aur.archlinux.org/) (AUR) can be used, which contains user made PKGBUILD files that are used to easily install programs. 

## Debian-Based

Debian is a point release distribution, meaning it is released in distinct versions, rather than constantly updated as new packages are available. This means you will never have the most up to date versions of programs, however it also means a more stable operating system which is less likely to break because of a package update. 

There are two main "types" of Debian distros. Pure Debian, and Ubuntu based. Ubuntu based distros will generally have *more* up to date packages. 

Debian uses [apt](https://en.wikipedia.org/wiki/APT_(software)) as a package manager, which contains a large number of packages. Note that Debian and Ubuntu have slightly different apt repositories, with Ubuntu generally having more packages available, but both are very usable. 

### Ubuntu

Standard [Ubuntu](https://ubuntu.com/) comes with the [GNOME](https://www.gnome.org/) desktop environment. This environment is the most similar environment to macOS. GNOME is designed to be very user friendly, and is one of the most common desktop environments to find in corporate or educational settings where Linux is in use (although not necessarily running atop Ubuntu).

### Kubuntu

[Kubuntu](https://kubuntu.org/) (and all distros ending in "buntu") is Ubuntu, but instead of GNOME, it comes with the [KDE](https://kde.org/) desktop environment. Often regarded as the most similar Linux DE to Windows, KDE will feel familiar to Windows users, while also being very customisable, arguably more so than GNOME. 

### Xubuntu

[Xubuntu](https://xubuntu.org/) (pronounced Zubuntu) comes with the [Xfce](xfce.md) desktop environment. Xfce is "Windows like" similar to KDE, however it is more lightweight and arguably more customisable. Xfce is more of a "minimal" desktop environment. 

### Lubuntu

[Lubuntu](https://lubuntu.me/) features the [LXQt](https://lxqt-project.org/) desktop environment. This is a lightweight window manager, making this distro excellent for older/slower hardware. It is also highly customisable, with LXQt Panel being similar to Xfce Panel in terms of customisability, and the underlying window manager [Openbox](http://openbox.org/wiki/Main_Page) also being very configurable. 

### Linux Mint - Ubuntu

[Linux Mint](https://www.linuxmint.com/download.php) by default is based on Ubuntu, but with the [Cinnamon](https://www.linuxmint.com/rel_victoria_cinnamon_whatsnew.php) desktop environment. Cinnamon is based on GNOME 3, and is somewhat of a cross between GNOME and KDE. As with KDE, it will feel familiar to Windows users. Unlike (U/Ku/Xu/Lu)buntu, Linux Mint does not come with Snap packages installed by default. 

### Linux Mint - Debian

Linux Mint also has a [Debian version](https://www.linuxmint.com/edition.php?id=308), which is based off of pure Debian instead of Ubuntu. 

### Debian

Much like Arch Linux, [Debian](https://www.debian.org/) allows you to install whichever desktop environment you choose during the setup process. If you want a **super stable** system, it is worth considering Debian, or if you like the sound of Ubuntu, but would rather avoid some of it's annoying quirks (such as the inclusion of Snap packages by default). The download and install process for Debian is more difficult than for any of the Ubuntu/Linux Mint distros, however still easier than Arch thanks to the included installer.
[Debian Install Guide - Chris Titus Tech](https://youtu.be/CJ41KZ0fBMc?si=gN_BxHZq-AjEusX_)

### Pop!_OS

[Pop!_OS](https://pop.system76.com/) is an Ubuntu based distro, using a modified version of the GNOME desktop environment. It comes with a number of enhancements, such as a tiling mode. Pop!_OS is often recommended for a first Linux distro, as it is very easy to setup and use, and comes with a good number of inbuilt features and programs. It is also good for gamers who want to try Linux, as there is a download that comes packaged with the NVidia drivers. 

## Fedora

[Fedora](https://fedoraproject.org/) is somewhat of a middle point between Arch and Ubuntu/Debian. If you like the sound of Ubuntu, but would prefer more up to date packages/kernel, as well as no annoying Snap packages, then Fedora is a good choice. Fedora is available with either [GNOME](https://fedoraproject.org/workstation/) or [KDE](https://fedoraproject.org/spins/kde/).

Although Fedora is more up to date than Ubuntu/Debian, it is still a point release distro. Fedora uses [DNF](https://en.wikipedia.org/wiki/DNF_(software)) as it's package manager.

## My Personal Recommendations

### Beginners

If you are a complete beginner to Linux, I would recommend going with an Ubuntu based distro, as they are the easiest to get beginner-level support with. As for which one, I think it depends which OS you are coming from. 

If you are coming from macOS, I would recommend Ubuntu or Pop!_OS for a GNOME desktop environment. GNOME not only looks the most similar to macOS, but the applications menu is also very similar to Apple's launchpad. Pop!_OS especially if you using an NVidia graphics card. 

If you are coming from Windows 8+, I would recommend Kubuntu for the KDE desktop environment, as the applications menu is very reminiscent of the Windows start menu, and the whole OS has a "Windows" feel to it IMO
  
If you are coming from an older version of Windows, or you are on a low powered PC, Xubuntu or Lubuntu are good options to try out, as the Xfce and LXQt desktop environments are very lightweight, and both reminiscent of older Windows versions (XP-7) in my opinion. 

### Gamers

The [Steam Deck](https://store.steampowered.com/steamdeck/), one of the main reasons for the uptick in interest in Linux gaming, runs on a version of Arch Linux, branded as SteamOS, meaning it has fully up to date packages/kernel. This is advantageous for gaming, as it will generally offer more performance, and gamers often have the latest hardware, which will be more supported in the up to date kernel. 

Unfortunately the [current public release of SteamOS](https://store.steampowered.com/steamos/buildyourown) is not this version, it's an older version based on Debian 8. If you feel confident installing Arch manually, I would recommend Arch for gaming, otherwise I would recommend Pop!_OS as it is very easy to setup, and comes with the NVidia drivers included, vastly simplifying driver setup. 

See [Gaming](gaming.md) for more details about gaming on Linux. 

### Customisation Nuts

If you are very into customising your OS and getting everything set up *exactly* as you like it, I would recommend Arch for the up to date packages, and Xfce as the Desktop Environment (unless you would prefer a tiling window manager). Arch allows you to build your system up from a very minimalist install, giving you more control over what is installed. Xfce as a desktop environment also allows a very high level of customisation, both in terms of theming and control over the wm (window movement, workspace management etc). It also has an excellent panel, allowing for a variety of custom widgets using the [Xfce Genmon Plugin](https://docs.xfce.org/panel-plugins/xfce4-genmon-plugin/start).
