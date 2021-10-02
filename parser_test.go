package astiffmpeg

import (
	"reflect"
	"testing"
	"time"

	"github.com/asticode/go-astikit"
)

func TestDefaultStdErrParser(t *testing.T) {
	p := &defaultStdErrParser{}
	e := DefaultStdErrResults{
		Bitrate: astikit.Float64Ptr(2065.5 * 1000),
		FPS:     astikit.IntPtr(254),
		Frame:   astikit.IntPtr(17448),
		Q:       astikit.Float64Ptr(31.0),
		Size:    astikit.IntPtr(176032 * 8 * 1000),
		Speed:   astikit.Float64Ptr(10.2),
		Time:    astikit.DurationPtr(140*time.Millisecond + 38*time.Second + 11*time.Minute),
	}
	g := p.parseResults([]byte("frame=17448 fps=254 q=31.0 size=  176032kB time=00:11:38.14 bitrate=2065.5kbits/s speed=10.2x"))
	if !reflect.DeepEqual(e, g) {
		t.Errorf("expected %+v, got %+v", e, g)
	}
}
