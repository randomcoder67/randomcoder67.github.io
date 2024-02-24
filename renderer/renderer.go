package main

import (
	"io"
	"github.com/gomarkdown/markdown/ast"
	"strconv"
	"fmt"
)

var _ = fmt.Println

var headings = []Heading{}
var inHeading bool = false
var headingLevel int = 1
var currentHeading []int = []int{0,0,0,0,0,0}
var contents = ""

type Heading struct {
	Level       int
	Title       string
	Subheadings []Heading
}

func renderContentsRecursive(curHeadings []Heading) {
	contents += "\n\t<ul>"
	for _, heading := range curHeadings {
		contents += "\n\t\t<li>"
		contents += "\n\t\t\t<a>"
		
		var levelString string = ""
		//fmt.Println("Headings:", curHeadings)
		currentHeading[heading.Level-1]++
		for i:=0; i<heading.Level; i++ {
			levelString = levelString + strconv.Itoa(currentHeading[i]) + "."
		}
		for i:=heading.Level; i<6; i++ {
			currentHeading[i] = 0
		}
		
		contents += "\n\t\t\t\t<span>" + levelString + "</span>"
		contents += "\n\t\t\t\t<span>" + heading.Title + "</span>"
		contents += "\n\t\t\t</a>"
		renderContentsRecursive(heading.Subheadings)
		contents += "\n\t\t</li>"
	} 
	contents += "\n\t</ul>"
}

func renderContents() string {
	contents = TOC_START
	
	currentHeading = []int{0,0,0,0,0,0}
	
	renderContentsRecursive(headings)
	
	_ = contents
	return contents
}

func renderInfoBox() string {
	return ""
}

func myRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if text, ok := node.(*ast.Link); ok {
		renderLinkHTMLExtension(w, text, entering)
		return ast.GoToNext, true
	}
	if heading, ok := node.(*ast.Heading); ok {
		renderHeading(w, heading, entering)
		return ast.GoToNext, true
	}
	if text, ok := node.(*ast.Text); ok {
		renderText(w, text)
		return ast.GoToNext, true
	}
	if imgLink, ok := node.(*ImgLink); ok {
		imgLinkRenderHook(w, imgLink, entering)
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}

// Function to render links, renders normally except changes .md extension to .html extension
func renderLinkHTMLExtension(w io.Writer, p *ast.Link, entering bool) {
	// Format: <a href="https://www.nasa.gov/multimedia/imagegallery/iotd.html" target="_blank">NASA Image of the Day</a>
	if entering {
		var link string = string(p.Destination)
		if len(link) > 0 && link[len(link)-3:] == ".md" {
			link = link[:len(link)-3] + ".html"
		}
		io.WriteString(w, "<a href=\"" + link + "\" target=\"_self\">")
	} else {
		io.WriteString(w, "</a>")
	}
}

// Function to render heading including the heading ID for contents
func renderHeading(w io.Writer, p *ast.Heading, entering bool) {
	headingLevel = p.Level
	level := strconv.Itoa(p.Level)
	if entering {
		//fmt.Println("BEFORE", currentHeading)
		currentHeading[headingLevel-1]++
		for i:=headingLevel; i<6; i++ {
			currentHeading[i] = 0
		}
		//fmt.Println("AFTER", currentHeading)
		io.WriteString(w, "<h" + level + " id=" + p.HeadingID + ">")
		inHeading = true
	} else {
		io.WriteString(w, "</h" + level + ">")
		inHeading = false
	}
}

// Text rendering function, works normally except if the text is a heading, in which case appends the heading to the heading array
func renderText(w io.Writer, p *ast.Text) {
	var textString string = string(p.Literal)
	io.WriteString(w, textString)
	if inHeading {
		//fmt.Printf("Headings Before: %+v\n", headings)
		//fmt.Println("Cur Pos:", currentHeading) 
		
		var workOn *[]Heading = &headings
		for i:=1; i<headingLevel; i++ {
			//fmt.Println("HERE", currentHeading[i-1], len((*workOn)))
			//fmt.Printf("Work On (%d): %+v\n", i, workOn)
			
			if currentHeading[i-1]-1 < 0 {
				*workOn = append(*workOn, Heading {
					Level: 0,
					Title: "",
				})
				currentHeading[i-1]++
			}
			
			workOn = &(*workOn)[currentHeading[i-1]-1].Subheadings
		}
		
		*workOn = append(*workOn, Heading {
			Level: headingLevel,
			Title: textString,
		})
		
		
		//fmt.Printf("Headings After: %+v\n\n", headings)
	}
}

func imgLinkRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if _, ok := node.(*ImgLink); ok {
		if entering {
			io.WriteString(w, "<ul id=\"linklist\">")
			
			for i, img := range node.(*ImgLink).ImageURLs {
				var altText string = node.(*ImgLink).AltText[i]
				var link string = node.(*ImgLink).Links[i]
				var fullString string = "<li><a href=\"" + link + "\"> <img src=\"" + img + "\"> <span class=\"title\">" + altText + "</span><span class=\"desc\">Test</span></a></li>"
				io.WriteString(w, fullString)
			}
			
			io.WriteString(w, "</ul>")
		}
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}
