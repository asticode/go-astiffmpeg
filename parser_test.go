package astiffmpeg

import (
	"testing"
	"time"

	"github.com/asticode/go-astitools/ptr"
	"github.com/stretchr/testify/assert"
)

func TestDefaultStdErrParser(t *testing.T) {
	p := &defaultStdErrParser{}
	r := p.parseResults([]byte("frame=17448 fps=254 q=31.0 size=  176032kB time=00:11:38.14 bitrate=2065.5kbits/s speed=10.2x"))
	assert.Equal(t, DefaultStdErrResults{
		Bitrate: astiptr.Float(2065.5 * 1000),
		FPS:     astiptr.Int(254),
		Frame:   astiptr.Int(17448),
		Q:       astiptr.Float(31.0),
		Size:    astiptr.Int(176032 * 8 * 1000),
		Speed:   astiptr.Float(10.2),
		Time:    astiptr.Duration(140*time.Millisecond + 38*time.Second + 11*time.Minute),
	}, r)
}
