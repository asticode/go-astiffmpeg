package astiffmpeg

import (
	"bytes"
	"strconv"
	"strings"
	"time"

	"github.com/asticode/go-astikit"
)

// StdErrParser represents an object capable of parsing stderr
type StdErrParser interface {
	Period() time.Duration
	Process(t time.Time, b *bytes.Buffer)
}

// DefaultStdErrParser creates the default stderr parser
func DefaultStdErrParser(period time.Duration, fn func(r DefaultStdErrResults)) StdErrParser {
	return &defaultStdErrParser{
		fn:     fn,
		period: period,
	}
}

type defaultStdErrParser struct {
	fn     func(r DefaultStdErrResults)
	period time.Duration
}

func (p defaultStdErrParser) Period() time.Duration {
	return p.period
}

func (p defaultStdErrParser) Process(t time.Time, b *bytes.Buffer) {
	// Split on \n
	var lines = bytes.Split(b.Bytes(), []byte("\n"))
	if len(lines) == 0 {
		return
	}

	// Split on \r
	var items = bytes.Split(lines[len(lines)-1], []byte("\r"))
	if len(items) < 2 {
		return
	}

	// Parse results
	r := p.parseResults(items[len(items)-2])

	// Execute callback
	p.fn(r)
}

// DefaultStdErrResults represents default stderr results
type DefaultStdErrResults struct {
	Bitrate *float64 // bits/s
	FPS     *int
	Frame   *int
	Q       *float64
	Size    *int // bits
	Speed   *float64
	Time    *time.Duration
}

// frame=17448 fps=254 q=31.0 size=  176032kB time=00:11:38.14 bitrate=2065.5kbits/s speed=10.2x
func (p defaultStdErrParser) parseResults(b []byte) (r DefaultStdErrResults) {
	// Split on =
	var k string
	for idx, i := range bytes.Split(b, []byte("=")) {
		// Split on space
		var items = bytes.Split(bytes.TrimSpace(i), []byte(" "))

		// Parse key/value
		if len(items[0]) > 0 && len(k) > 0 {
			v := string(items[0])
			switch k {
			case "bitrate":
				// There may be other suffix, but we only support this one for now
				v = strings.TrimSuffix(v, "kbits/s")
				if p, err := strconv.ParseFloat(v, 64); err == nil {
					r.Bitrate = astikit.Float64Ptr(p * 1000)
				}
			case "frame":
				if p, err := strconv.Atoi(v); err == nil {
					r.Frame = astikit.IntPtr(p)
				}
			case "fps":
				if p, err := strconv.ParseFloat(v, 64); err == nil {
					r.FPS = astikit.IntPtr(int(p))
				}
			case "q":
				if p, err := strconv.ParseFloat(v, 64); err == nil {
					r.Q = astikit.Float64Ptr(p)
				}
			case "size":
				if n, err := numberFromString(v); err == nil {
					r.Size = astikit.IntPtr(int(n.float64()))
				}
			case "speed":
				if p, err := strconv.ParseFloat(strings.TrimSuffix(v, "x"), 64); err == nil {
					r.Speed = astikit.Float64Ptr(p)
				}
			case "time":
				// Split on .
				var d time.Duration
				ps := strings.Split(v, ".")
				if len(ps) > 1 {
					if p, err := strconv.Atoi(ps[1]); err == nil {
						// For now we make the assumption that milliseconds are in this format ".99" and not ".999"
						d += time.Duration(p*10) * time.Millisecond
					}
				}

				// Split on :
				ps = strings.Split(ps[0], ":")
				if len(ps) >= 3 {
					if p, err := strconv.Atoi(ps[0]); err == nil {
						d += time.Duration(p) * time.Hour
					}
					if p, err := strconv.Atoi(ps[1]); err == nil {
						d += time.Duration(p) * time.Minute
					}
					if p, err := strconv.Atoi(ps[2]); err == nil {
						d += time.Duration(p) * time.Second
					}
				}
				r.Time = astikit.DurationPtr(d)
			}
		}

		// Get key
		if len(items) > 1 && len(items[1]) > 0 {
			k = string(items[1])
		} else if len(items) == 1 && idx == 0 {
			k = string(items[0])
		}
	}
	return
}
