package main

import (
	"github.com/gomarkdown/markdown/ast"
	"bytes"
	"strings"
	"fmt"
)

var _ = fmt.Println

type CustomDiv struct {
	ast.Leaf
	End      bool
	Class    string
}

type Tags struct {
	ast.Leaf
}

type InfoBox struct {
	ast.Leaf
	Title   string
	Image   string
	Caption string
	Data    [][]string
}

var tags = []byte("\\tags ")
var infobox = []byte("\\infobox\n")
var customDiv = []byte("\\customDiv ")

func parseVarious(data []byte) (ast.Node, []byte, int) {
	if bytes.HasPrefix(data, infobox) {
		return parseInfobox(data)
	} else if bytes.HasPrefix(data, customDiv) {
		return parseCustomDiv(data)
	} else if bytes.HasPrefix(data, tags) {
		return parseTags(data)
	}
	return nil, nil, 0
}

func parseCustomDiv(data []byte) (ast.Node, []byte, int) {
	i := len(customDiv)
	end := bytes.Index(data[i:], []byte("\n\n"))
	if end < 0 {
		return nil, data, 0
	}
	end = end + i

	split := strings.Split(string(data[i:end]), " ")
	res := &CustomDiv{}
	
	if len(split) == 1 {
		res.End = true
	} else {
		res.End = false
		res.Class = split[1]
	}
	
	return res, nil, end
}

func parseInfobox(data []byte) (ast.Node, []byte, int) {
	i := len(infobox)
	end := bytes.Index(data[i:], []byte("\n\n"))
	if end < 0 {
		return nil, data, 0
	}
	end = end + i

	split := strings.Split(string(data[i:end]), "\n")
	res := &InfoBox{}

	for _, line := range split {
		if line[:6] == "title:" {
			res.Title = line[7:]
		} else if line [:4] == "img:" {
			res.Image = line[5:]
		} else if line [:8] == "caption:" {
			res.Caption = line[9:]
		} else {
			parts := strings.Split(line, "|")
			res.Data = append(res.Data, []string{parts[0], parts[1]})
		}
	}
	
	return res, nil, end
}

func parseTags(data []byte) (ast.Node, []byte, int) {
	i := len(infobox)
	end := bytes.Index(data[i:], []byte("\n"))
	if end < 0 {
		return nil, data, 0
	}
	end = end + i

	res := &Tags{}
	return res, nil, end
}

