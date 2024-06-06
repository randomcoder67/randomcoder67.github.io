package main

import (
	"os"
	"fmt"
	"strings"
	"encoding/json"
)

const CSS_FILE_LOCATION = "/renderer/style.css"

type Options struct {
	Folder        bool
	InputName     string
	OutputName    string
	ForceOverwrite bool
}

/*
func parseArgs(args []string) Options {
	var i int = 0
	var curOptions Options = Options {}
	
	for i<len(args) {
		switch args[i] {
		case "-w":
			curOptions.WikiFormat = true
			i++
		case "-f":
			curOptions.Folder = true
			i++
		default:
			if curOptions.InputName == "" {
				curOptions.InputName = args[i]
			} else {
				curOptions.OutputName = args[i]
			}
			i++
		}
	}
	
	if curOptions.OutputName == "" {
		if curOptions.Folder {
			errorOut()
		}
		
		var inputName string = curOptions.InputName
		var lenInput int = len(inputName)
		curOptions.OutputName = inputName[:lenInput-3] + ".html"
	}
	
	return curOptions
}
*/

func autoFormatOutputName(opt Options) string {
	var inputName string = opt.InputName
	var lenInput int = len(inputName)
	return inputName[:lenInput-3] + ".html"
}

func parseArgs(args []string) Options {
	var opt Options = Options {}

	if len(args) == 1 {
		opt.InputName = args[0]
		opt.OutputName = autoFormatOutputName(opt)
	} else if len(args) == 2 {
		opt.InputName = args[0]
		opt.OutputName = args[1]
	} else if len(args) == 3 {
		if args[0][0:2] == "-f" {
			if args[0] == "-ff" {
				opt.ForceOverwrite = true
			} else {
				opt.ForceOverwrite = false
			}
			opt.Folder = true
			opt.InputName = args[1]
			opt.OutputName = args[2]
		} else {
			errorOut()
		}
	} else {
		errorOut()
	}
	
	return opt
}

func errorOut() {
	fmt.Println("Error, incorrectly formatted arguments")
	printHelp()
	os.Exit(1)
}

func printHelp() {
	fmt.Printf("Usage:\n  render input.md [output.html] (defaults to input with .html extension)\n  render -f[f] inputDir/ outputDir/ (second f to force existing folder overwrite)\n")
}

func main() {
	opt := parseArgs(os.Args[1:])
	//fmt.Printf("%+v\n", opt)

	contentRootDir = opt.InputName
	lastSlash := strings.LastIndex(opt.InputName, "/")
	if lastSlash == len(opt.InputName) - 1 {
		lastSlash = strings.LastIndex(opt.InputName[:len(opt.InputName)-1], "/")
	}
	if lastSlash == -1 {
		rootDir = "."
	} else {
		rootDir = opt.InputName[:lastSlash]
	}
	
	if opt.Folder {
		renderFolder(opt.InputName, opt.OutputName, opt.ForceOverwrite)
	} else {
		renderFile(opt.InputName, opt.OutputName, 0)
	}
	//fmt.Println("BACK IN MAIN\n\n")
	_ = json.MarshalIndent
	//thing, _ := json.MarshalIndent(headings, "", "  ")
	//fmt.Println(string(thing))
	//fmt.Printf("%+v\n", headings)
	//thing := renderContents()
	//fmt.Println(thing)
}
