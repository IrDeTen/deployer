package models

import (
	"strings"

	"github.com/Delta456/box-cli-maker"
)

type VBox struct {
	Box      box.Box
	Title    string
	LineSize int
}

func (vb *VBox) Print(inpText string) {
	var (
		outLines []string = make([]string, 0)
	)

	if vb.LineSize == 0 {
		vb.Box.Print(vb.Title, inpText)
		return
	}
	inpLines := strings.Split(inpText, "\n")
	for _, line := range inpLines {
		outLines = append(outLines, vb.linePrepare(line))
	}
	vb.Box.Println(vb.Title, strings.Join(outLines, "\n"))
}

func (vb *VBox) linePrepare(inpLine string) string {
	var (
		temp         string
		words, parts []string
	)

	inpLN := len(inpLine)
	if inpLN <= vb.LineSize {
		return vb.buildUP(inpLine)
	}

	words = strings.Split(inpLine, " ")
	for i, val := range words {
		length := len(temp)
		if length+len(val) > vb.LineSize {
			parts = append(parts, vb.buildUP(temp))
			temp = val
			continue
		}
		if i == len(words)-1 {
			parts = append(parts, vb.buildUP(val))
			break
		}
		if len(temp) != 0 {
			temp = temp + " "
		}
		temp = temp + val
	}

	return strings.Join(parts, "\n")
}

func (vb *VBox) buildUP(inp string) string {
	inp = inp + strings.Repeat(" ", vb.LineSize-len(inp))
	return inp
}
