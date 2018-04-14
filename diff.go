package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/adam-hanna/arrayOperations"
	"github.com/pmezard/go-difflib/difflib"
)

var wg sync.WaitGroup

// const srca = "D:/code/igv/nrithai/BR-GV-FOR-THAGMO-PROD-REL-15-20032018/istar-gv/devel/gv-deployments/gv-th/gv-th-inf/"
// const srcb = "D:/code/igv/nrithai/BR-GV-FOR-THAGMO-PROD-REL-16-27032018/istar-gv/devel/gv-deployments/gv-th/gv-th-inf/"

const srca = "D:/code/igv/nrithai/BR-GV-FOR-THAGMO-PROD-REL-15-20032018/istar-gv/devel/gv-core/"
const srcb = "D:/code/igv/nrithai/BR-GV-FOR-THAGMO-PROD-REL-16-27032018/istar-gv/devel/gv-core/"

var fh filesHolder

func main() {

	start := time.Now()

	fh = New()

	fileNamesA := identifyFiles(srca, fh)
	fileNamesB := identifyFiles(srcb, fh)

	unique := arrayOperations.IntersectString(fileNamesA, fileNamesB)
	//	uniqueA := arrayOperations.DifferenceString(fileNamesA, fileNamesB)
	//	uniqueB := arrayOperations.DifferenceString(fileNamesB, fileNamesA)

	// copyFileNamesA := make([]string, len(fileNamesA))
	// copy(copyFileNamesA, fileNamesA)

	// copyFileNamesB := make([]string, len(fileNamesB))
	// copy(copyFileNamesB, fileNamesB)

	/*
		unique := make([]string, min(len(fileNamesA), len(fileNamesB)))
		uniquelen := 0

		uniqueAOnly := make([]string, 0)

		//remove present files in a

		for i := range fileNamesA {
			for j := range fileNamesB {

				if fileNamesA[i] == fileNamesB[j] {
					// copyFileNamesA[i] = ""
					// copyFileNamesB[j] = ""
					unique[uniquelen] = fileNamesA[i]
					uniquelen++
					break
				}
				//if it has come here then A has this and B doesnt
				uniqueAOnly = append(uniqueAOnly, fileNamesA[i])
			}
		}

		//delete unique extras
		unique = unique[:uniquelen]
	*/
	for i := range unique {
		go findDiff(unique[i])
		wg.Add(1)
	}

	wg.Wait()

	end := time.Now()
	fmt.Println("Time taken : ", end.Sub(start))

}

func findDiff(newpath string) {

	defer wg.Done()

	data, _ := ioutil.ReadFile(srca + newpath)
	data2, _ := ioutil.ReadFile(srcb + newpath)

	var newfilecontentA []byte
	var newfilecontentB []byte

	newfilecontentA = fh.removeComments(newpath, data)
	newfilecontentB = fh.removeComments(newpath, data2)

	fname := newpath
	idx := strings.LastIndex(newpath, "/")
	if idx == -1 {
		idx = strings.LastIndex(newpath, "\\")
	}
	if idx != -1 {
		fname = newpath[idx+1:]
	}

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

		old := make(map[int]*patchList)

		splitOutput := difflib.SplitLines(text)

		for i := 0; i < len(splitOutput); i++ {
			x := splitOutput[i]
			if strings.HasPrefix(x, "@@") {
				lineNo, lineCount := extractOldLineNo(x)
				old[lineNo] = lineCount

			}
		} //

		//print old ones
		oldText := ""

		for i := 0; i < len(splitLinesA); i++ {
			lNo := i + 1
			if old[lNo] != nil {
				for j := 0; j < old[lNo].oldDeleteCount; j++ {
					oldText += "<del>" + html.EscapeString(splitLinesA[i+j]) + "</del>" + "<br/>"
				}
				i = i + old[lNo].oldDeleteCount

				//add the new count
				for z := 0; z < old[lNo].newAddCount; z++ {

					oldText += "<ins>" + html.EscapeString(splitLinesB[z+old[lNo].newLineNo-1]) + "</ins>" + "<br/>"
				}

			} else {
				oldText += html.EscapeString(splitLinesA[i]) + "<br/>"
			}

		}

		txt := "<html><head> <style type=\"text/css\"> ins{background-color:green;} del{background-color:red;} </style> </head><body>"
		txt += "<h1>" + fname + "</h1>"
		txt += oldText
		txt += "</body></html>"

		ioutil.WriteFile("d:/tmp/"+fname+".html", []byte(txt), 0755)

	} // end diff

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

// func processDiffWithActualFileContent(data []byte, diffs []diffmatchpatch.Diff) {
// 	//loop over each and every stupid line just for fun.. and then print based on Equal / Delete or something
// 	//currCount := 0
// 	for _, diff := range diffs {
// 		switch diff.Type.String() {

// 		case "Insert":

// 		case "Delete":

// 		case "Equal":
// 			//print all lines but calling it equal

// 		}

// 	}

// }

func saveFile(newpath string, data string) {
	lastIdx := strings.LastIndex(newpath, "\\")
	filename := newpath
	if lastIdx != -1 {
		filename = newpath[lastIdx:]
	}

	fmt.Println("d:/tmp/" + filename + ".html")
	ioutil.WriteFile("d:/tmp/"+filename+".html", []byte(data), 0755)
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
