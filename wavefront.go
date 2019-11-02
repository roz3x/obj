package obj

//Wavefront the file handler
type Wavefront struct {
	Name             string
	CoordinatePoints [][]float64
	LineIndices      [][]int
	VertexNormals    [][]float64
	Faces            [][][]int
}
