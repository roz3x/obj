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
	lines             [][]byte
	coordinatesPoints [][]float64
	lineindexs        [][]int
	vecternormal      [][]float64
	faces             [][][]int
}

//Init function that does all the work
func (f *Wavefront) Init(file io.Reader) {

	var (
		linesregex       = regexp.MustCompile(`.*\n`)
		pointlinetest    = regexp.MustCompile(`^v\ `)
		pointtest        = regexp.MustCompile(`[\d]*\.[\d]*`)
		linelinetest     = regexp.MustCompile(`^l`)
		linetest         = regexp.MustCompile(`[\d]+`)
		vecterNormalLine = regexp.MustCompile(`^vn`)
		facelinetest     = regexp.MustCompile(`^f`)
		faceblocktest    = regexp.MustCompile(`[\d|/]+`)
		facesettest      = regexp.MustCompile(`[\d]+`)
	)
	filebody, err := ioutil.ReadAll(file)
	handle(err)
	// this line will remove error related to last missing newline character
	filebody = append(filebody, []byte("\n")...)
	f.lines = linesregex.FindAll(filebody, -1)
	for _, curLine := range f.lines {
		if pointlinetest.Match(curLine) {
			tmp := []float64{}
			for _, k := range pointtest.FindAll(curLine, -1) {
				tmp = append(tmp, convFloat(strconv.ParseFloat(string(k), 64)))
			}
			f.coordinatesPoints = append(f.coordinatesPoints, tmp)
		} else if linelinetest.Match(curLine) {
			tmp := []int{}
			for _, k := range linetest.FindAll(curLine, -1) {
				tmp = append(tmp, convInt(strconv.ParseInt(string(k), 10, 64)))
			}
			f.lineindexs = append(f.lineindexs, tmp)
		} else if vecterNormalLine.Match(curLine) {
			tmp := []float64{}
			for _, k := range pointtest.FindAll(curLine, -1) {
				tmp = append(tmp, convFloat(strconv.ParseFloat(string(k), 64)))
			}
			f.vecternormal = append(f.vecternormal, tmp)
		} else if facelinetest.Match(curLine) {
			tmp := faceblocktest.FindAll(curLine, -1)
			tmpfarr := [][]int{}
			for _, i := range tmp {
				p := facesettest.FindAll(i, -1)
				tmparr := []int{}
				for _, k := range p {
					tmparr = append(tmparr, convInt(strconv.ParseInt(string(k), 10, 64)))
				}
				tmpfarr = append(tmpfarr, tmparr)
			}
			f.faces = append(f.faces, tmpfarr)
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
