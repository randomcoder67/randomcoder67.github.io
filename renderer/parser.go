package main

import (
	"github.com/gomarkdown/markdown/ast"
	"bytes"
	"strings"
	"fmt"
	"encoding/json"
)

var _ = fmt.Println

type CustomDiv struct {
	ast.Leaf
	End      bool
	Class    string
}

type ImgLink struct {
	ast.Leaf
	ImageURLs []string
	TitleText []string
	HoverText []string
	Links     []string
	Size      string
	Align     string
}

type ImgLinkOptions struct {
	Size  string `json:"size"`
	Align string `json:"align"`
}

var imgLink = []byte("\\imglink\n")
var infobox = []byte("\\infobox\n")
var customDiv = []byte("\\customDiv ")

func parseVarious(data []byte) (ast.Node, []byte, int) {
	if bytes.HasPrefix(data, imgLink) {
		return parseImgLink(data)
	} else if bytes.HasPrefix(data, infobox) {
		return parseInfobox(data)
	} else if bytes.HasPrefix(data, customDiv) {
		return parseCustomDiv(data)
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

func parseImgLink(data []byte) (ast.Node, []byte, int) {
	i := len(imgLink)
	end := bytes.Index(data[i:], []byte("\n\n"))
	if end < 0 {
		return nil, data, 0
	}
	
	end = end + i
	
	lines := string(data[i:end])
	parts := strings.Split(lines, "\n")
	
	var align string = "left"
	var size string = "100%"
	var contentLoc int = 0
	
	
	if parts[0][0] == '{' {
		contentLoc = 1
		var options ImgLinkOptions
		err := json.Unmarshal([]byte(parts[0]), &options)
		if err != nil {
			panic(err)
		}
		if options.Size != "" {
			size = options.Size
		}
		if options.Align != "" {
			align = options.Align
		}
	}
	
	res := &ImgLink {
		Align: align,
		Size: size,
	}
	
	for i:=contentLoc; i<len(parts); i++ {
		bits := strings.Split(parts[i], "|")
		res.ImageURLs = append(res.ImageURLs, bits[0])
		res.TitleText = append(res.TitleText, bits[1])
		res.HoverText = append(res.HoverText, bits[2])
		res.Links = append(res.Links, bits[3])
	}
	
	//fmt.Printf("%+v\n", res)
	
	return res, nil, end
}

func parseInfobox(data []byte) (ast.Node, []byte, int) {
	res := &CustomDiv{
	}
	var end int = 0
	return res, nil, end
}

