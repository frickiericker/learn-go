package matrix

import (
	"testing"
)

func TestNewSquare_checkSize(t *testing.T) {
	sq := NewSquare(7)
	if result := sq.Size(); result != 7 {
		t.Errorf("unexpected result from Size() %v", result)
	}
	if result := sq.Rows(); result != 7 {
		t.Errorf("unexpected result from Rows() %v", result)
	}
	if result := sq.Cols(); result != 7 {
		t.Errorf("unexpected result from Cols() %v", result)
	}
	if result := sq.Data(); len(result) != 7*7 {
		t.Errorf("unexpected result from Data() %v", result)
	}
}

func TestSquare_Get_safeRange(t *testing.T) {
	sq := NewSquare(5)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			sq.Get(i, j)
		}
	}
}

func TestSquare_Set_safeRange(t *testing.T) {
	sq := NewSquare(5)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			sq.Set(i, j, 1.23)
		}
	}
}
