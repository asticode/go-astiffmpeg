package astiffmpeg

import (
	"math"
	"testing"
)

func TestNumber(t *testing.T) {
	for _, i := range []struct {
		f        float64
		hasError bool
		i        string
		s        string
	}{
		{hasError: true, i: "12abc"},
		{f: 12.0, i: "12", s: "12"},
		{f: 12.0 * 8, i: "12B", s: "12B"},
		{f: 12.0 * 8 * 1000, i: "12kB", s: "12kB"},
		{f: 12.0 * 8 * 1024, i: "12kiB", s: "12kiB"},
		{f: 12.0 * 8 * math.Pow(1024, 2), i: "12MiB", s: "12MiB"},
		{f: 12.0 * 8 * math.Pow(1024, 3), i: "12GiB", s: "12GiB"},
		{f: 12.0 * 8 * math.Pow(1024, 4), i: "12TiB", s: "12TiB"},
		{f: 12.0 * 8 * math.Pow(1024, 5), i: "12PiB", s: "12PiB"},
	} {
		n, err := numberFromString(i.i)
		if i.hasError {
			if err == nil {
				t.Error("expected error")
			}
		} else {
			if err != nil {
				t.Errorf("expected no error, got %s", err.Error())
			}
			if i.f != n.float64() {
				t.Errorf("expected %v, got %v", i.f, n.float64())
			}
			if i.s != n.string() {
				t.Errorf("expected %s, got %s", i.s, n.string())
			}
		}
	}
}
