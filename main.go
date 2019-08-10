package main

import (
	"fmt"
	"os"
)

type wavefront struct {
	lines             [][]byte
	coordinatesPoints [][]float64
	lineindexs        [][]int
	vecternormal      [][]float64
	faces             [][][]int
}

func main() {

	f, err := os.Open("sample.obj")

	handle(err)

	newWave := &wavefront{}

	newWave.init(f)

	fmt.Printf("%v  %v %v  \n", newWave.faces, newWave.lineindexs, newWave.vecternormal)

}
