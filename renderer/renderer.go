package main

import (
	"io"
	"github.com/gomarkdown/markdown/ast"
	"strconv"
	"strings"
	"fmt"
	"os"
	"image"
	_ "golang.org/x/image/webp"
	_ "image/jpeg"
	_ "image/png"
)

var _ = fmt.Println

var headings = []Heading{}
var inHeading bool = false
var inImage bool = false
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
		var link string = strings.ToLower(heading.Title)
		link = strings.ReplaceAll(link, " ", "-")
		link = strings.ReplaceAll(link, "--", "")
		link = strings.ReplaceAll(link, "/", "-")
		contents += "\n\t\t\t<a href=\"#" + link + "\">"
		
		var levelString string = ""
		
		currentHeading[heading.Level-1]++
		for i:=1; i<heading.Level; i++ {
			levelString = levelString + strconv.Itoa(currentHeading[i]) + "."
		}
		levelString = levelString[:len(levelString)-1]
		for i:=heading.Level; i<6; i++ {
			currentHeading[i] = 0
		}
		
		contents += "\n\t\t\t\t<span class=\"tocnumber\">" + levelString + "</span>"
		contents += "\n\t\t\t\t<span class=\"toctext\">" + heading.Title + "</span>"
		contents += "\n\t\t\t</a>"
		renderContentsRecursive(heading.Subheadings)
		contents += "\n\t\t</li>"
	} 
	contents += "\n\t</ul>"
}

func renderContents() string {
	contents = TOC_START
	
	currentHeading = []int{0,0,0,0,0,0}
	
	renderContentsRecursive(headings[0].Subheadings)
	
	contents = contents + TOC_END
	return contents
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
	if customDiv, ok := node.(*CustomDiv); ok {
		renderCustomDiv(w, customDiv, entering)
		return ast.GoToNext, true
	}
	if infobox, ok := node.(*InfoBox); ok {
		renderInfoBox(w, infobox, entering)
		return ast.GoToNext, true
	}
	if img, ok := node.(*ast.Image); ok {
		imgRenderHook(w, img, entering)
		return ast.GoToNext, true
	}
	if _, ok := node.(*Tags); ok {
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}

func getImageDimensions(link string) (int, int) {
	var fullLink string

	if link[0] == '/' {
		fullLink = contentRootDir + "/" + link
	} else {
		fullLink = currentDir + "/" + link
	}
	
	reader, err := os.Open(fullLink)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	
	im, _, err := image.DecodeConfig(reader)
	if err != nil {
		panic(err)
	}
	return im.Width, im.Height
}


func imgRenderHook(w io.Writer, p *ast.Image, entering bool) {
	if entering {
		inImage = true
		var altText string = string(p.Children[0].AsLeaf().Literal)
		var link string = string(p.Destination)
		
		width, height := getImageDimensions(link)
		
		fmt.Fprintf(w, "<img src=\"%s\" alt=\"%s\" width=\"%d\" height=\"%d\">\n", link, altText, width, height)
	} else {
		inImage = false
	}
}

// Function to render links, renders normally except changes .md extension to .html extension
func renderLinkHTMLExtension(w io.Writer, p *ast.Link, entering bool) {
	// Format: <a href="https://www.nasa.gov/multimedia/imagegallery/iotd.html" target="_blank">NASA Image of the Day</a>
	if entering {
		var link string = string(p.Destination)
		if len(link) > 0 && link[len(link)-3:] == ".md" {
			link = link[:len(link)-3] + ".html"
		} else if strings.Contains(link, ".md#") {
			link = strings.ReplaceAll(link, ".md#", ".html#")
		}
		io.WriteString(w, "<a href=\"" + link + "\" target=\"_self\">")
	} else {
		io.WriteString(w, "</a>\n")
	}
}

// Function to render heading including the heading ID for contents
func renderHeading(w io.Writer, p *ast.Heading, entering bool) {
	headingLevel = p.Level
	level := strconv.Itoa(p.Level)
	if entering {
		currentHeading[headingLevel-1]++
		for i:=headingLevel; i<6; i++ {
			currentHeading[i] = 0
		}
		if headingLevel != 1 {
			io.WriteString(w, "<h" + level + " id=" + p.HeadingID + ">")
		}
		inHeading = true
	} else {
		if headingLevel != 1 {
			io.WriteString(w, "</h" + level + ">\n")
		}
		inHeading = false
	}
}

// Text rendering function, works normally except if the text is a heading, in which case appends the heading to the heading array
func renderText(w io.Writer, p *ast.Text) {
	if inImage {
		return
	}
	var textString string = string(p.Literal)
	if inHeading && headingLevel != 1 {
		io.WriteString(w, textString)
	}
	if !inHeading {
		io.WriteString(w, textString)
	}
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

func renderCustomDiv(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if _, ok := node.(*CustomDiv); ok {
		if entering {
			data, _ := node.(*CustomDiv)
			if data.End {
				io.WriteString(w, "</div>\n")
			} else {
				io.WriteString(w, "<div class=\"" + data.Class + "\">")
			}
		}
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}

func renderInfoBox(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if _, ok := node.(*InfoBox); ok {
		if entering {
			data, _ := node.(*InfoBox)
			io.WriteString(w, "<table class=\"infobox\">\n")
			io.WriteString(w, "<caption>" + data.Title + "</caption>\n")
			io.WriteString(w, "<tbody>\n")


			if data.Image != "" {
				width, height := getImageDimensions(data.Image)
				
				io.WriteString(w, "<tr>\n")
				io.WriteString(w, "<td class=\"image_caption\" colspan=\"2\">\n")
				io.WriteString(w, "<img alt=\"" + data.Caption + "\" src=\"" + data.Image + "\" width=\"" + strconv.Itoa(width) + "\" width=\"" + strconv.Itoa(height) + "\">\n")
				io.WriteString(w, "<div>" + data.Caption + "</div>\n")
				io.WriteString(w, "</td>\n")
				io.WriteString(w, "</tr>\n")
			}

			for _, dataPoint := range data.Data {
				io.WriteString(w, "<tr>\n")
				io.WriteString(w, "<td class=\"item\">" + dataPoint[0] + "</td>\n")
				var content string = dataPoint[1]
				if content[0] == '[' {
					var text string = strings.Split(content[1:], "]")[0]
					var link string = strings.Split(strings.Split(content, "(")[1], ")")[0]
					content = fmt.Sprintf("<a href=\"%s\">%s</a>", link, text)
				}
				io.WriteString(w, "<td>" + content + "</td>\n")
				io.WriteString(w, "</tr>\n")
			}

			io.WriteString(w, "</tbody>\n")
			io.WriteString(w, "</table>\n")
		}
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}
