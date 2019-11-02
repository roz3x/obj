package main

import (
	"github.com/roz3x/obj"
	"fmt"
)
//the defination of Wavefornt struct 
//use this to manipulate the data
//type Wavefront struct {
//	Name             string
//	CoordinatePoints [][]float64
//	LineIndices      [][]int
//	VertexNormals    [][]float64
//	Faces            [][][]int
//}

func main() {
	wfd := &obj.Wavefront{}
	wfs := &obj.Wavefront{}
	wft := &obj.Wavefront{}
	wfs.Decode("source.obj")
	wfd.Decode("destination.obj")
	obj.Copy(wfs,wfd,wft)
	wft.Encode("output.obj")
}
