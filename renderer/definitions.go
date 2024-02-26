package main

const TOC_START string = `<div id="toc" class="toc">
	<div id="toctitle" class="toctitle">
		<h2>Contents</h2>
		<span class="toctoggle">
			[
			<a class="togglelink" role="button" tabindex="0">hide</a>
			]
		</span>
	</div>`

const TOC_END string = "</div>"

const TITLE_START string = `<!DOCTYPE html>
<html>
<head>
    <link rel="stylesheet" href="`

const TITLE_MIDDLE = `style.css">
<title>`

const TITLE_END string = `</title>
</head>`

const INFOBOX_START string = `<table class="infobox" style="width:22em">
	<caption class="summary">`

const INFOBOX_MIDDLE string = `</caption>
	<tbody>`

const INFOBOX_END string = `	</tbody>
</table>`

const NAV_BAR string = `<div class="topnav">
    <a href="/index.html" class="title">randomcoder67</a>
    <a href="https://youtube.com" class="youtubeLogo"><img src="/logos/youtubeFullColour.png" alt="YouTube Logo" width="40" height="auto" class="youtubeLogo"></a>
    <a href="https://github.com/randomcoder67" class="githubLogo"><img src="/logos/github-mark-white.png" alt="GitHub Logo" width="32" height="auto" class="githubLogo"></a>
    <a href="/about.html" class="text">About</a>
    <a href="/wiki/index.html" class="text">Wiki</a>
    <a href="#" class="text">Travel</a>
    <a href="#" class="text">Cycling</a>
    <a href="#" class="text">Guides</a>
    <a href="myprograms/landing.html" class="text">Coding</a>   
</div>`

const CONTENT_START = "<div class=\"content\">"
const CONTENT_END = "</div>"

/*
		<tr>
			<td colspan=2>
				<img alt="wanker" src="images/mediaui.png" width="300">
			</td>
		</tr>
		<tr>
			<th>Initial Release</th>
			<td>May 31st 2020</td>
		</tr>
		<tr>
			<th>Stable Release</th>
			<td>June 29th 2020</td>
		</tr>
		<tr>
			<th>Language</th>
			<td>C and Golang</td>
		</tr>
		<tr>
			<th>Engine</th>
			<td>Unity and Shite</td>
		</tr>
		<tr>
			<th>Operating System</th>
			<td>Linux, Windows, macOS, Android</td>
		</tr>
		<tr>
			<th>License</th>
			<td>None mate fuck off</td>
		</tr>
*/

/*
	<ul>
		<li class="toclevel-1 tocsection-1">
			<a href="#heading-1">
				<span class="tocnumber">1</span>
				<span class="toctext">In-game description</span>
			</a>
		</li>
		<li clas="toclevel-1 tocsection-2">
			<a href="#heading-2">
				<span class="tocnumber">2</span>
				<span class="toctext">Atmosphere</span>
			</a>
			<ul>
				<li class="toclevel-2 tocsection-3">
					<a href="#heading-2-1">
						<span class="tocnumber">2.1</span>
						<span class="toctext">Wanking off</span>
					</a>
				</li>
			</ul>
		</li>
	</ul>
*/
