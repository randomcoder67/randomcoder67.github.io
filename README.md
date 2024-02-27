# randomcoder67.github.io

My personal website (blog and wiki)

I decided to make a website, for now I am hosting it on GitHub pages. Eventually I might move to a custom domain name and VPS hosting.  
The purpose of the website is to serve as a blog, as well as a wiki for various bits of random knowledge I have gathered (mostly so I don't forget stuff)

## Editing

If you spot an issue or something missing in the wiki, feel free to submit and issue or pull request. The website is edited as markdown

## Rendering

The site is rendered using [gomarkdown](https://pkg.go.dev/github.com/gomarkdown/markdown), with some custom parsing to handle infoboxes, link-img lists etc. I will document the rendering scripts on the Wiki at some point.  
There are also various helper functions which help in managing the content, for example managing tags, and getting a list of the most recent pages.

## Credit

* The aforementioned [gomarkdown](https://pkg.go.dev/github.com/gomarkdown/markdown) is used to render the markdown to HTML
* HTML converted from Markdown with: [gomarkdown](https://pkg.go.dev/github.com/gomarkdown/markdown)  
* Website inspired and CSS based off of [Luke Smith's Website](https://lukesmith.xyz)  
* Further inspiration for website and wiki part from [LinuxReviews](https://linuxreviews.org/LinuxReviews)  
* Website hosted with [GitHub Pages](https://pages.github.com/)  
