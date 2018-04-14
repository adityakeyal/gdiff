package main

type fileParser interface {
	getExtn() string
	considerFile(filename string) bool
	removeComments(filecontent []byte) []byte
}
