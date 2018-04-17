package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
)

const (
	a = 1
	e = 0
	d = -1
)

type mergedFileData struct {
	ops  int
	line string
}

func (x mergedFileData) toHTML(lineNo int) string {

	text := "<tr>"

	line := strings.Replace(x.line, "\n", "", -1)
	line = strings.Replace(line, "\r", "", -1)

	switch x.ops {
	case a:
		text += `  <td class="d2h-code-linenumber d2h-ins d2h-change"> <div class="line-num1">` + strconv.Itoa(lineNo) + `</div> <div class="line-num2"></div> </td>`
		text += `<td class="d2h-ins d2h-change"> <div class="d2h-code-line d2h-ins d2h-change"> <span class="d2h-code-line-prefix">+</span> <span class="d2h-code-line-ctn">  ` + html.EscapeString(line) + ` </span> </div> </td>`
	case e:
		text += `  <td class="d2h-code-linenumber d2h-ctx"> <div class="line-num1">` + strconv.Itoa(lineNo) + `</div> <div class="line-num2"></div> </td>`
		text += `<td class="d2h-cntx"> <div class="d2h-code-line d2h-cntx"> <span class="d2h-code-line-prefix"> </span> <span class="d2h-code-line-ctn">` + html.EscapeString(line) + `</span> </div>  </td>`
	case d:
		text += `  <td class="d2h-code-linenumber d2h-del d2h-change"> <div class="line-num1"></div> <div class="line-num2"></div> </td>`
		text += `<td class="d2h-del d2h-change"> <div class="d2h-code-line d2h-del d2h-change"><span class="d2h-code-line-prefix">-</span><span class="d2h-code-line-ctn"> ` + html.EscapeString(line) + `</span></div></td>`
	}

	text += "</tr>"
	return text
}

func processDiff(fd fileDiff) []mergedFileData {

	//fmt.Println(fd)
	old := make(map[int]*patchList)

	splitOutput := difflib.SplitLines(fd.diffText)

	for i := 0; i < len(splitOutput); i++ {
		x := splitOutput[i]
		if strings.HasPrefix(x, "@@") {
			lineNo, lineCount := extractOldLineNo(x)
			old[lineNo] = lineCount

		}
	} //

	//create a file array with a length of 0 and capacity of old file size (approx)
	finalFile := make([]mergedFileData, 0, len(fd.linesFileA))

	for i := 0; i < len(fd.linesFileA); i++ {
		lNo := i + 1
		if old[lNo] != nil {
			for j := 0; j < old[lNo].oldDeleteCount; j++ {
				//oldText += "<del>" + html.EscapeString() + "</del>" + "<br/>"
				finalFile = append(finalFile, mergedFileData{-1, fd.linesFileA[i+j]})
			}
			i = i + old[lNo].oldDeleteCount

			//add the new count
			for z := 0; z < old[lNo].newAddCount; z++ {
				finalFile = append(finalFile, mergedFileData{+1, fd.linesFileB[z+old[lNo].newLineNo-1]})
				//oldText += "<ins>" + html.EscapeString() + "</ins>" + "<br/>"
			}

		} else {
			finalFile = append(finalFile, mergedFileData{0, fd.linesFileA[i]})
			//oldText += html.EscapeString() + "<br/>"
		}

	}

	return finalFile

	// txt := "<html><head> <style type=\"text/css\"> ins{background-color:green;} del{background-color:red;} </style> </head><body>"
	// txt += "<h1>" + fd.fileName + "</h1>"
	// txt += oldText
	// txt += "</body></html>"

	// ioutil.WriteFile("d:/tmp/"+fd.fileName+".html", []byte(txt), 0755)
}

func extractOldLineNo(x string) (int, *patchList) {
	old := strings.Split(x, " ")[1]
	vals := strings.Split(old, ",")

	lineNo, _ := strconv.Atoi(vals[0][1:])
	pp := new(patchList)

	if len(vals) == 1 {
		pp.oldDeleteCount = 1
	} else {
		pp.oldDeleteCount, _ = strconv.Atoi(vals[1])
	}

	pp.newLineNo, pp.newAddCount = extractNewLineNo(x)

	return lineNo, pp

}

func extractNewLineNo(x string) (int, int) {
	old := strings.Split(x, " ")[2]
	vals := strings.Split(old, ",")

	if len(vals) == 1 {
		lineNo, _ := strconv.Atoi(vals[0])
		return lineNo, 1
	}

	lineNo, _ := strconv.Atoi(vals[0])
	count, _ := strconv.Atoi(vals[1])

	return lineNo, count

}

func saveFile(newpath string, data string) {
	lastIdx := strings.LastIndex(newpath, "\\")
	filename := newpath
	if lastIdx != -1 {
		filename = newpath[lastIdx:]
	}

	fmt.Println("d:/tmp/" + filename + ".html")
	ioutil.WriteFile("d:/tmp/"+filename+".html", []byte(data), 0755)
}
