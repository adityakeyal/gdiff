package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
)

func findDiff(newpath string) {

	defer wg.Done()

	data, _ := ioutil.ReadFile(srca + newpath)
	data2, _ := ioutil.ReadFile(srcb + newpath)

	var newfilecontentA []byte
	var newfilecontentB []byte

	newfilecontentA = fh.removeComments(newpath, data)
	newfilecontentB = fh.removeComments(newpath, data2)

	fileNameArray := strings.Split(strings.Replace(newpath, "\\", "/", -1), "/")

	fname := ""
	maxLength := min(3, len(fileNameArray))

	for i := maxLength; i > 0; i-- {
		fname += fileNameArray[len(fileNameArray)-i]
		if i > 1 {
			fname += "_"
		}
	}

	// for (x := maxLength; x>=0; x--) {
	// 	fmt.Println(i)
	// }

	// idx := strings.LastIndex(newpath, "/")
	// if idx == -1 {
	// 	idx = strings.LastIndex(newpath, "\\")
	// }
	// if idx != -1 {
	// 	fname = newpath[idx+1:]
	// }

	//Option 2

	splitLinesA := difflib.SplitLines(string(data))
	splitLinesB := difflib.SplitLines(string(data2))

	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(string(newfilecontentA)),
		B:        difflib.SplitLines(string(newfilecontentB)),
		FromFile: "Original " + fname,
		ToFile:   "Current " + fname,
		//Context:  3,
	}
	text, _ := difflib.GetUnifiedDiffString(diff)

	if text != "" {
		fd := fileDiff{fname, text, splitLinesA, splitLinesB}
		finalFile := processDiff(fd)
		printHtml(fname, finalFile)

	} // end diff

}

type filesHolder struct {
	parsers []fileParser
}

func (f filesHolder) isValidExtn(path string) bool {
	for i := 0; i < len(f.parsers); i++ {
		if f.parsers[i].considerFile(path) {
			return true
		}
	}
	return false
}

func (f filesHolder) removeComments(fname string, content []byte) []byte {
	for i := 0; i < len(f.parsers); i++ {
		if strings.HasSuffix(fname, f.parsers[i].getExtn()) {
			return f.parsers[i].removeComments(content)
		}
	}
	return nil
}

func New() filesHolder {
	var f filesHolder
	var j fileParser = javaFileType{}
	var x fileParser = xmlFileType{}
	var p fileParser = propFileType{}
	f.parsers = []fileParser{j, x, p}
	return f
}

type patchList struct {
	oldDeleteCount int
	newLineNo      int
	newAddCount    int
}

type fileDiff struct {
	fileName, diffText     string
	linesFileA, linesFileB []string
}

///////////////////////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////////////

func printHtml(fileName string, finalFile []mergedFileData) {

	f, err := os.Create("d:\\tmp\\" + fileName + ".html")
	defer f.Close()
	check(err)

	_, err = f.WriteString("<html><head> ")
	check(err)
	_, err = f.WriteString(printHeadStyle)
	check(err)
	_, err = f.WriteString("</head><body>")
	check(err)

	_, err = f.WriteString(`<div class="d2h-file-header">
		<span class="d2h-file-name-wrapper">
		<span class="d2h-icon-wrapper"><svg aria-hidden="true" class="d2h-icon" height="16" version="1.1" viewBox="0 0 12 16" width="12">
		<path d="M6 5H2v-1h4v1zM2 8h7v-1H2v1z m0 2h7v-1H2v1z m0 2h7v-1H2v1z m10-7.5v9.5c0 0.55-0.45 1-1 1H1c-0.55 0-1-0.45-1-1V2c0-0.55 0.45-1 1-1h7.5l3.5 3.5z m-1 0.5L8 2H1v12h10V5z"></path>
		</svg></span>
		<span class="d2h-file-name">` + fileName + `</span>
		</div>`)
	check(err)

	_, err = f.WriteString(`<div class="d2h-file-diff"><div class="d2h-code-wrapper"> <table class="d2h-diff-table selecting-right">  <tbody class="d2h-diff-tbody">`)

	for i := 0; i < (len(finalFile)); i++ {

		x := finalFile[i]
		_, err = f.WriteString(x.toHTML(i + 1))
		check(err)

	}

	_, err = f.WriteString("</tbody></table></div></div>")
	_, err = f.WriteString("</body></html>")
	//	f.Sync()
	// txt := "<html><head> <style type=\"text/css\"> ins{background-color:green;} del{background-color:red;} </style> </head><body>"
	// txt += "<h1>" + fd.fileName + "</h1>"
	// txt += oldText
	// txt += "</body></html>"

	// ioutil.WriteFile("d:/tmp/"+fd.fileName+".html", []byte(txt), 0755)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
