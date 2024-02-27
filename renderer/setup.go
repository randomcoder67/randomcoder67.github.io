package main

import (
	"os"
	"io"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	TYPE_FILE int = 0
	TYPE_DIR int = 1
	TYPE_NOTHING int = 2
)

// Check if a given path points to a file, a directory, or nothing
func isFileOrDir(filename string) int {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return TYPE_NOTHING
	} else if fileInfo.IsDir() {
		return TYPE_DIR
	}
	return TYPE_FILE
}

func parserHook(data []byte) (ast.Node, []byte, int) {
	if node, d, n := parseVarious(data); node != nil {
		return node, d, n
	}
	return nil, nil, 0
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	p.Opts.ParserHook = parserHook
	doc := p.Parse(md)
	

	// create HTML renderer with extensions
	//htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions {
		Flags: html.CommonFlags | html.HrefTargetBlank,
		RenderNodeHook: myRenderHook,
	}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func renderFile(inputFile string, outputFile string, level int) {
	headings = []Heading{}
	currentHeading = []int{0,0,0,0,0,0}

	// Open file
	fmt.Println("Rendering:", inputFile)
	mds, _ := os.ReadFile(inputFile)
	// Convert markdown to html
	md := []byte(mds)
	html := mdToHTML(md)
	
	// Write html to file
	//var cssInclude string = "<link rel=\"stylesheet\" type=\"text/css\" href=\"" + strings.Repeat("../", level) + "style.css\"/>"
	
	var finalHTML strings.Builder
	finalHTML.WriteString(TITLE_START)
	finalHTML.WriteString(strings.Repeat("../", level))
	finalHTML.WriteString(TITLE_MIDDLE)
	finalHTML.WriteString(headings[0].Title)
	finalHTML.WriteString(TITLE_END)
	finalHTML.WriteString("\n")
	finalHTML.WriteString(NAV_BAR)
	finalHTML.WriteString("\n")
	finalHTML.WriteString(CONTENT_START)
	finalHTML.WriteString("\n")
	finalHTML.WriteString(string(html))
	finalHTML.WriteString(CONTENT_END)
	finalHTML.WriteString("\n")

	var htmlString string = finalHTML.String()
	if strings.Contains(outputFile, "/wiki/") {
		var contents string = renderContents()
		oldString := htmlString
		htmlString = strings.Replace(htmlString, "<h2", contents + "\n<h2", 1)
	}
	
	err := os.WriteFile(outputFile, []byte(htmlString), 0666)
	if err != nil {
		panic(err)
	}
}

func renderFolder(dirName string, outputDir string) {
	// Check dirName is a directory
	if isFileOrDir(dirName) != TYPE_DIR {
		fmt.Println("Error, not a directory")
		os.Exit(1)
	}

	
	// Check outputDir has a slash
	if outputDir[len(outputDir)-1] != '/' {
		outputDir = outputDir + "/"
	}

	// Check inputDir has a slash
	if dirName[len(dirName)-1] != '/' {
		dirName = dirName + "/"
	}
	// Make the "new root" directory
	if isFileOrDir(outputDir) != TYPE_DIR {
		os.Mkdir(outputDir, os.ModePerm)
	}
	// Copy style.css there
	homeDir, _ := os.UserHomeDir()
	_ = homeDir
	source, err := os.Open(CSS_FILE_LOCATION)
	if err != nil {
		panic(err)
	}
	// Open destination
	destination, err := os.Create(outputDir + "style.css")
	if err != nil {
		panic(err)
	}
	// Copy source to destination
	_, err = io.Copy(destination, source)
	if err != nil {
		panic(err)
	}
	
	// Recursive function to go through given folder and subfolders
	var listInDir func (string, int)
	listInDir = func(dirNameCur string, level int) {
		//fmt.Println(level)
		files, err := ioutil.ReadDir(dirNameCur)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			if file.IsDir() { // If file is a directory, make it in new location and go a level deeper
				if isFileOrDir(strings.ReplaceAll(outputDir + dirNameCur + "/" + file.Name(), dirName, "")) != TYPE_DIR {
					err := os.Mkdir(strings.ReplaceAll(outputDir + dirNameCur + "/" + file.Name(), dirName, ""), os.ModePerm)
					if err != nil {
						panic(err)
					}
				}
				listInDir(dirNameCur + "/" + file.Name(), level+1)
			} else { // If file is a file, add it to allFiles
				// Get new file name
				var newFileName string = strings.ReplaceAll(outputDir + dirNameCur + "/" + file.Name(), dirName, "")
				// If file is a markdown file
				if newFileName[len(newFileName)-3:] == ".md" {
					// Change extension to html
					newFileName = newFileName[:len(newFileName)-3] + ".html"
					//fmt.Println(dirNameCur + "/" + file.Name(), newFileName)
					// Render file
					renderFile(dirNameCur + "/" + file.Name(), newFileName, level)
				} else { // If file is not a markdown file, just copy
					// Open source
					source, err := os.Open(dirNameCur + "/" + file.Name())
					if err != nil {
						panic(err)
					}
					// Open destination
					destination, err := os.Create(newFileName)
					if err != nil {
						panic(err)
					}
					// Copy source to destination
					_, err = io.Copy(destination, source)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
	listInDir(dirName[:len(dirName)-1], 0)
}
