package main

import (
	"strings"
)

/*

Comments can be of 3 varieties

a. // which will remove all the elements after the // (with the exception of being present inside a ")
b. multi line comments
c. multi line comments in same line

*/

const newLine byte = '\n'

type javaFileType struct {
}

func (j javaFileType) considerFile(filename string) bool {

	return (strings.HasSuffix(filename, j.getExtn()))
}

func (j javaFileType) removeComments(filecontent []byte) []byte {

	b := make([]byte, len(filecontent))

	countOfBytes := 0

	isSingleLineComment, isMultiLineComment := false, false

	for i := 0; i < len(filecontent); i++ {
		x := filecontent[i]

		//check for single line comment
		if (i+1) < len(filecontent) && x == '/' && filecontent[i+1] == '/' {
			isSingleLineComment = true
		}

		if isSingleLineComment && x == newLine {
			isSingleLineComment = false
		}

		//check for multi line comment

		if (i+1) < len(filecontent) && x == '/' && filecontent[i+1] == '*' {
			isMultiLineComment = true
		}

		if isMultiLineComment && i > 2 && filecontent[i-2] == '*' && filecontent[i-1] == '/' {
			isMultiLineComment = false

		}

		if (!isSingleLineComment && !isMultiLineComment) || x == '\n' {
			if i < len(b) {

				b[countOfBytes] = x
				countOfBytes++
			}
		}

	}

	if countOfBytes > 0 {
		return b[0 : countOfBytes-1]
	}

	var empty = []byte{' '}

	return empty

}

func (j javaFileType) getExtn() string {
	return ".java"
}
