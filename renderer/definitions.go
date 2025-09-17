package main

/*
const TOC_START string = `<div id="toc" class="toc">
	<div id="toctitle" class="toctitle">
		<h2>Contents</h2>
		<span class="toctoggle">
			[
			<a class="togglelink" role="button" tabindex="0">hide</a>
			]
		</span>
	</div>`
*/

const TOC_START string = `<div id="toc" class="toc">
	<div id="toctitle" class="toctitle">
		<h2>Contents</h2>
	</div>`

const TOC_END string = "</div>"

const TITLE_START string = `<!DOCTYPE html>
<html>
<head>
    <link rel="stylesheet" href="`

const TITLE_MIDDLE = `/style.css">
<title>`

const TITLE_END string = `</title>
</head>
<body>
<div class="page-wrapper">`

const INFOBOX_START string = `<table class="infobox" style="width:22em">
	<caption class="summary">`

const INFOBOX_MIDDLE string = `</caption>
	<tbody>`

const INFOBOX_END string = `	</tbody>
</table>`

const NAV_BAR string = `<div class="left-pane">
     <div class="left-title">
      <a href="/">Welcome To randomcoder67's Personal Website</a>
     </div>
     <div class="left-bar">
      <ul>
        <li><a href="/wiki/index.html">Wiki</a></li>
        <li><a href="/blog/index.html">Blog</a></li>
        <li><a href="/reviews/index.html">Reviews</a></li>
        <li><a href="/projects/index.html">Projects</a></li>
        <li><a href="/photography/index.html">Photography</a></li>
        <li><a href="/other/index.html">Other</a></li>
        <li><a href="/about.html">About</a></li>
        <li><a href="/contributions.html">Contributions</a></li>
        <li><a href="/about.html#contact">Contact</a></li>
      </ul>
     </div>
     <div class="left-bar-two">
      <ul>
        <li><a href="https://github.com/randomcoder67">GitHub</a></li>
        <li><a href="https://codeberg.org/randomcoder67">Codeberg</a></li>
        <li><a href="">YouTube</a></li>
        <li><a href="https://bsky.app/profile/randomcoder67.bsky.social">BlueSky</a></li>
        <li><a href="">RSS</a></li>
      </ul>
     </div>
    </div>`

const HEADING_ONE_START = `<div class="main-content">
<div class="center-div">
<div class="hanging-sign">
<h1 id="`

const HEADING_ONE_MIDDLE = `">`

const HEADING_ONE_END = `</h1>
</div>
</div>`

const CONTENT_START = "<div class=\"actual-content\">"
const CONTENT_END = "</div>"

const FOOTER = `</div><footer>footer test</footer>`

const PAGE_END = `</div>
</body>
</html>`
