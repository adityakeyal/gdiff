package main

import "strings"

/*

Comments can be of 1 variety

a. #

*/

type propFileType struct {
}

func (p propFileType) removeComments(filecontent []byte) []byte {

	b := make([]byte, len(filecontent))

	countOfBytes := 0

	isComment := false

	for i := 0; i < len(filecontent); i++ {
		x := filecontent[i]

		//check for single line comment
		if filecontent[i] == '#' {
			isComment = true
		}

		if isComment && filecontent[i] == '\n' {
			isComment = false
		}

		if !isComment {
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

func (p propFileType) getExtn() string {
	return ".properties"
}

func (p propFileType) considerFile(filename string) bool {

	if strings.HasSuffix(filename, p.getExtn()) {
		return strings.Index(filename, "/target/") == -1 && strings.Index(filename, "\\target\\") == -1
	}

	return false

}
