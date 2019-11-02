package obj

import (
	"fmt"
	"os"
)

func (f *Wavefront) Encode(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if f.Name == "" {
		fmt.Fprintf(file, "o default\n")
	} else {
		fmt.Fprintf(file, "o %v\n", f.Name)
	}
	for _, points := range f.CoordinatePoints {
		fmt.Fprintf(file, "v %v %v %v\n", points[0], points[1], points[2])
	}
	for _, normals := range f.VertexNormals {
		fmt.Fprintf(file, "vn ")
		for normal := range normals {
			fmt.Fprintf(file, "%v ", normals[normal])
		}
		fmt.Fprintf(file, "\n")
	}
	for _, faces := range f.Faces {
		fmt.Fprintf(file, "f ")
		for _, sets := range faces {
			fmt.Fprintf(file, " %v// ", sets[0])
		}
		fmt.Fprintf(file, "\n")
	}
	for _, lines := range f.LineIndices {
		fmt.Fprintf(file, "l %v %v\n", lines[0], lines[1])
	}

}
