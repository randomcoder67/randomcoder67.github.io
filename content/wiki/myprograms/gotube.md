# GoTube Terminal YouTube Client

\infobox
title: GoTube
img: /images/gotube.png
caption: GoTube Example Search
Creater|randomcoder67
GitHub Link|[GoTube Repo](https://github.com/randomcoder67/GoTube-Terminal-YouTube)
Language|Golang
Version|Beta 0.2.0
Category|TUI Application

GoTube is a terminal YouTube client, designed to emulate as much of the browser YouTube experience as possible, without the slow website, and without the ads. Similarly to [yt-dlp](https://github.com/yt-dlp/yt-dlp), GoTube makes use of the saved browser cookies to allow access to account features. This means that creating an API key and going through the OAuth process is not necessary, and also allows GoTube access to certain functions that the official API does not have, such as viewing and editing your Watch Later playing, and saving video watch times to your history.

GoTube was designed to be used side by side with browser YouTube, or the mobile app. For this reason, GoTube operates on real YouTube playlists, instead of locally saved playlists, and is capabale of both resuming videos from their previous watch point, and saving a videos watch time to your history, allowing you to pick up on another PC, or on mobile.  

[Ueberzug](https://github.com/ueber-devel/ueberzug) is used to allow thumbnails to be rendered in any* terminal emulator, with the main display mode emulating the YouTube grid interface.

## Design

GoTube does not use the [official YouTube API](https://developers.google.com/youtube/v3), instead using the website API, through use of browser cookies. There are two main parts to this:

1. The browser cookies - The relevant data is extracted from Firefox cookies to retrieve YouTube webpages. This means that the subscription feed will be accurate for the users account, searches will be saved, recommendations will be accurate etc.
2. The "SAPISIDHASH". This is a code used in API calls made by YouTube. It is a combination of data from a specific cookie, as well as the current [Unix timestamp](https://en.wikipedia.org/wiki/Unix_time), hashed with the [SHA1](https://en.wikipedia.org/wiki/SHA-1) algorithm. This code allows interacting with a user's account features, such as creating and deleting playlists, adding items to playlists etc

The TUI interface for GoTube is written using [Tcell V2](https://github.com/gdamore/tcell), a low level Golang TUI library. Ueberzug is used to display images.

## Usage

In order to use GoTube, you must be logged into YouTube in Firefox.

The TUI can be navigated using the arrow keys or vim keys.

## How it Works

### Cookies

### SAPISIDHASH

### Syncing Video Watchtime

### Ueberzug

### Internal Logic
