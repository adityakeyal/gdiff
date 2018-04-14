package main

import "strings"

/*

Comments can be of 1 variety

a. <!-- -->

*/

type xmlFileType struct {
}

func (x xmlFileType) removeComments(filecontent []byte) []byte {

	b := make([]byte, len(filecontent))

	countOfBytes := 0

	isComment := false

	for i := 0; i < len(filecontent); i++ {
		x := filecontent[i]

		//check for single line comment
		if (i+3) < len(filecontent) && filecontent[i] == '<' && filecontent[i+1] == '!' && filecontent[i+2] == '-' && filecontent[i+3] == '-' {
			isComment = true
		}

		if !isComment {
			if i < len(b) {
				b[countOfBytes] = x
				countOfBytes++
			}
		}

		//check for multi line comment

		if isComment && i > 2 && filecontent[i-2] == '-' && filecontent[i-1] == '-' && filecontent[i] == '>' {
			isComment = false

		}

	}

	if countOfBytes > 0 {
		return b[0 : countOfBytes-1]
	}

	var empty = []byte{' '}
	return empty

}

func (x xmlFileType) getExtn() string {
	return ".xml"
}

func (x xmlFileType) considerFile(filename string) bool {

	if strings.HasSuffix(filename, x.getExtn()) {
		return strings.Index(filename, "/target/") == -1 && strings.Index(filename, "\\target\\") == -1
	}

	return false
}
