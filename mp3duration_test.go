package mp3duration

import "testing"

func TestCbr(t *testing.T) {
	duration, _ := Calculate("testdata/demo - cbr.mp3")
	ms := duration.Milliseconds()
	if ms != 285779 {
		t.Errorf("FAILED: Expected 285779 got %v", ms)
	} else {
		t.Log("PASSED")
	}
}

func TestVbr(t *testing.T) {
	duration, _ := Calculate("testdata/demo - vbr.mp3")
	ms := duration.Milliseconds()
	if ms != 285727 {
		t.Errorf("FAILED: Expected 285727 got %v", ms)
	} else {
		t.Log("PASSED")
	}
}
