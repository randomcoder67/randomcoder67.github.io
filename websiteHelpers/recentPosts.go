package main

import (
	"fmt"
	"os/exec"
	"os"
	"time"
	"strings"
)

type Entry struct {
	Title string
	Link  string
	Date  time.Time
}

var numWanted int

var mostRecent = []Entry{}

func fileExists(filename string) bool {
	_, err := os.Stat("../" + filename)
	if err != nil {
		return false
	}
	return true
}

func resolveNewFilename(oldFilename string) string {
	var newFilename string = oldFilename
	cmd := exec.Command("git", "log", "--format=format:%s", "--name-status", "--diff-filter=R")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(output), "\n")
	var movedLines = []string{}
	for i:=len(lines)-1; i>0; i-- {
		var line string = lines[i]
		if len(line) < 1 {
			continue
		}
		if line[0] == 'R' {
			movedLines = append(movedLines, line)
		}
	}

	for _, line := range movedLines {
		parts := strings.Fields(line)
		if parts[1] == newFilename {
			newFilename = parts[2]
		}
	}
	return newFilename
}

func getFileTitle(filename string) string {
	filename = "../" + filename
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	
	firstLine := strings.Split(string(dat), "\n")[0]
	return firstLine[2:]
}

func getLatest() {
	cmd := exec.Command("git", "log", "--format=format:%ci", "--name-only", "--diff-filter=A")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	commits := strings.Split(string(output), "\n\n")
	
	for _, dateRange := range commits {
		var curDate string = strings.Split(strings.Split(dateRange, "\n")[0], " ")[0]
		lines := strings.Split(dateRange, "\n")

		for _, line := range lines {
			if line == "README.md" {
				continue
			}
			if len(line) < 4 {
				continue
			}
			if line[len(line)-3:] == ".md" {
				parsedDate, err := time.Parse("2006-01-02", curDate)
				if err != nil {
					panic(err)
				}
				if !fileExists(line) {
					line = resolveNewFilename(line)
				}
				
				newEntry := Entry{
					Title: getFileTitle(line),
					Link: line[8:],
					Date: parsedDate,
				}
				mostRecent = append(mostRecent, newEntry)
				if len(mostRecent) == numWanted {
					return
				}
			}
		}
	}
}

func renderList() string {
	var sb strings.Builder
	for _, entry := range mostRecent {
		sb.WriteString("* ")
		sb.WriteString(entry.Date.Format("Mon 02 Jan 2006"))
		sb.WriteString(" - ")
		sb.WriteString("[")
		sb.WriteString(entry.Title)
		sb.WriteString("](")
		sb.WriteString(entry.Link)
		sb.WriteString(")\n")
	}
	return sb.String()
}

func GetNewestArticles(numToGet int) {
	numWanted = numToGet
	getLatest()
	var rendered string = renderList()
	fmt.Printf(rendered)
}
