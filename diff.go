package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/adam-hanna/arrayOperations"
)

// const srca = "D:/code/igv/nrithai/BR-GV-FOR-THAGMO-PROD-REL-15-20032018/istar-gv/devel/gv-deployments/gv-th/gv-th-inf/"
// const srcb = "D:/code/igv/nrithai/BR-GV-FOR-THAGMO-PROD-REL-16-27032018/istar-gv/devel/gv-deployments/gv-th/gv-th-inf/"

const srca = "D:/code/igv/nrithai/BR-GV-FOR-THAGMO-PROD-REL-15-20032018/istar-gv/devel/gv-core/"
const srcb = "D:/code/igv/nrithai/BR-GV-FOR-THAGMO-PROD-REL-16-27032018/istar-gv/devel/gv-core/"

var wg sync.WaitGroup
var fh filesHolder

func main() {
	start := time.Now()

	fh = New()

	fileNamesA := identifyFiles(srca, fh)
	fileNamesB := identifyFiles(srcb, fh)

	unique := arrayOperations.IntersectString(fileNamesA, fileNamesB)

	for i := range unique {
		go findDiff(unique[i])
		wg.Add(1)
	}

	wg.Wait()

	end := time.Now()
	fmt.Println("Time taken : ", end.Sub(start))

}
