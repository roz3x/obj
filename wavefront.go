package obj

import (
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

//Wavefront the file handler
type Wavefront struct {
	Lines            [][]byte
	CoordinatePoints [][]float64
	LineIndices      [][]int
	VertexNormals    [][]float64
	Faces            [][][]int
}

//Init function that does all the work
func (f *Wavefront) Init(file io.Reader) {

	var (
		Linesregex        = regexp.MustCompile(`.*\n`)
		pointlinetest     = regexp.MustCompile(`^v\ `)
		pointtest         = regexp.MustCompile(`[\d]*\.[\d]*`)
		linelinetest      = regexp.MustCompile(`^l`)
		linetest          = regexp.MustCompile(`[\d]+`)
		VertexNormalsLine = regexp.MustCompile(`^vn`)
		facelinetest      = regexp.MustCompile(`^f`)
		faceblocktest     = regexp.MustCompile(`[\d|/]+`)
		Facesettest       = regexp.MustCompile(`[\d]+`)
	)
	filebody, err := ioutil.ReadAll(file)
	handle(err)
	// this line will remove error related to last missing newline character
	filebody = append(filebody, []byte("\n")...)
	f.Lines = Linesregex.FindAll(filebody, -1)
	for _, curLine := range f.Lines {
		if pointlinetest.Match(curLine) {
			tmp := []float64{}
			for _, k := range pointtest.FindAll(curLine, -1) {
				tmp = append(tmp, convFloat(strconv.ParseFloat(string(k), 64)))
			}
			f.CoordinatePoints = append(f.CoordinatePoints, tmp)
		} else if linelinetest.Match(curLine) {
			tmp := []int{}
			for _, k := range linetest.FindAll(curLine, -1) {
				tmp = append(tmp, convInt(strconv.ParseInt(string(k), 10, 64)))
			}
			f.LineIndices = append(f.LineIndices, tmp)
		} else if VertexNormalsLine.Match(curLine) {
			tmp := []float64{}
			for _, k := range pointtest.FindAll(curLine, -1) {
				tmp = append(tmp, convFloat(strconv.ParseFloat(string(k), 64)))
			}
			f.VertexNormals = append(f.VertexNormals, tmp)
		} else if facelinetest.Match(curLine) {
			tmp := faceblocktest.FindAll(curLine, -1)
			tmpfarr := [][]int{}
			for _, i := range tmp {
				p := Facesettest.FindAll(i, -1)
				tmparr := []int{}
				for _, k := range p {
					tmparr = append(tmparr, convInt(strconv.ParseInt(string(k), 10, 64)))
				}
				tmpfarr = append(tmpfarr, tmparr)
			}
			f.Faces = append(f.Faces, tmpfarr)
		}
	}
}

func handle(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func convFloat(f float64, e error) float64 {
	handle(e)
	return f
}

func convInt(f int64, e error) int {
	handle(e)
	return int(f)

}
