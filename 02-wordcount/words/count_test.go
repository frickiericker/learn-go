package words

import (
	"testing"
)

func TestCount_emptyInput(t *testing.T) {
	if count := Count(""); count != 0 {
		t.Errorf("%v", count)
	}
}

func TestCount_one(t *testing.T) {
	if count := Count("foo"); count != 1 {
		t.Errorf("%v", count)
	}
}

func TestCount_two(t *testing.T) {
	if count := Count("apple banana"); count != 2 {
		t.Errorf("%v", count)
	}
}

func TestCount_singleLetter(t *testing.T) {
	if count := Count("x"); count != 1 {
		t.Errorf("%v", count)
	}
}

func TestCount_startsWithSpace(t *testing.T) {
	if count := Count(" foo bar baz"); count != 3 {
		t.Errorf("%v", count)
	}
}

func TestCount_endsWithSpace(t *testing.T) {
	if count := Count("foo bar baz "); count != 3 {
		t.Errorf("%v", count)
	}
}

func TestCount_punctuationChars(t *testing.T) {
	if count := Count("foo, bar, baz"); count != 3 {
		t.Errorf("%v", count)
	}
}

func TestCount_symbols(t *testing.T) {
	if count := Count("(x - y) / z * w"); count != 4 {
		t.Errorf("%v", count)
	}
}
