# yt-dlp

yt-dlp is command line YouTube downloader written in Python 3. A fork of the original youtube-dl, yt-dlp comes with significant improvements in performance over it's predecessor. As well as downloading YouTube videos, yt-dlp can be used to retrieve information about videos or playlists, as well as accessing information from an individuals YouTube account given the correct cookies. 

## Installation

### Arch

`pacman -S yt-dlp`

## Downloading Videos

`yt-dlp url` will download the given video in the best video and audio quality avalible. When used with YouTube, the downloaded file will generally be a .webm containter, with VP9 encoded video and Opus encoded audio. For most purposes, this is okay, as most modern GPUs include hardware VP9 Decode. ([Intel has since Skylake for example](https://en.wikipedia.org/wiki/Intel_Quick_Sync_Video#Hardware_decoding_and_encoding))  
However for phones or older PCs, H264 may be desired. This is covered below in the video format section. 

### Video Format

`yt-dlp -F url` is used to list the available video formats. This will return a table with entries similar to the following:

![yt-dlp -F output example](images/ytdlpOutput.png)

The columns are labelled at the top, and most are self-explanatory.  
The second column shows the container the video will be stored in and the eighth column shows the codec the video is encoded in. Note VP9 is `vp09`, H264 is `avc1` and [AV1](https://en.wikipedia.org/wiki/AV1) is `av01`.  
For audio, the codec will appear in the tenth column, options being Opus or M4A. Which one you use is essentially entirely dependant on whether you use mp4 or webm as a container.  
It is worth pointing out that most of the formats displayed will be either audio *or* video, not both. This is due to the way YouTube delivers content. 

`yt-dlp -f format url` is used to set video format. `format` refers to the numbers in the left most column of the `yt-dlp -F` output. Multiple formats can be selected at once, for example to combine audio and video into one. This is done using `+` to connect them. For example `yt-dlp -f 248+251 url`. 

### Extra Inclusions

`--embed-chapters` will include video chapters in the output file  
`--cookies-from-browser browser` will attempt to use cookies from the given browser, this is useful to download private or login-restricted videos  
`--cookies file.ext` is similar but will load cookies from a specific file

### Output File Name and Location

The default output template is: `%(title)s [%(id)s].%(ext)s`. This results in filenames including the YouTube video ID. In order to disable this, 

### Video Information

Do not extract the videos of a playlist, only list them 

`--flat-playlist` is an option used to stop the parsing of individual videos, and rather just gives information about the playlist, and limited information about the videos. 
