package seq

import (
	"reflect"
	"testing"
)

func TestFloat64Slice_Len(t *testing.T) {
	seq := Float64Slice{[]float64{1, 2, 3, 4}}
	if seq.Len() != 4 {
		t.Errorf("%v", seq)
	}
}

func TestFloat64Slice_Swap(t *testing.T) {
	seq := Float64Slice{[]float64{1, 2, 3, 4}}
	seq.Swap(1, 2)
	if !reflect.DeepEqual(seq.slice, []float64{1, 3, 2, 4}) {
		t.Errorf("%v", seq)
	}
}
