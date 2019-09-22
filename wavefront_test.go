package obj

import (
	"os"
	"testing"
)


//function now exposed 
func TestInit(t *testing.T) {
	f, err := os.Open("sample.obj")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	ob := &Wavefront{}
	ob.Init(f)
	if len(ob.CoordinatePoints) == 0 {
		t.Fail()
	}
}
