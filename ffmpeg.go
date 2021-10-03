package astiffmpeg

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// FFMpeg represents an entity capable of running an FFMpeg binary
// https://ffmpeg.org/ffmpeg.html
type FFMpeg struct {
	binaryPath   string
	stdErrParser StdErrParser
}

// New creates a new FFMpeg
func New(c Configuration) *FFMpeg {
	return &FFMpeg{binaryPath: c.BinaryPath}
}

// SetStdErrParser sets the stderr parser
func (f *FFMpeg) SetStdErrParser(s StdErrParser) {
	f.stdErrParser = s
}

// Exec executes the binary with the specified options
// ffmpeg [global_options] {[input_file_options] -i input_url} ... [output_file_options] output_url
func (f *FFMpeg) Exec(ctx context.Context, g GlobalOptions, in []Input, out Output) (err error) {
	// Create cmd
	var cmd = exec.CommandContext(ctx, f.binaryPath)
	cmd.Env = os.Environ()

	// Output is redirected in stderr only
	var bufErr = &bytes.Buffer{}
	cmd.Stderr = bufErr

	// Global options
	g.adaptCmd(cmd)

	// Parse stderr
	if f.stdErrParser != nil {
		t := time.NewTicker(f.stdErrParser.Period())
		defer t.Stop()
		go func() {
			for t := range t.C {
				f.stdErrParser.Process(t, bufErr)
			}
		}()
	}

	// Inputs
	for idx, i := range in {
		if err = i.adaptCmd(cmd); err != nil {
			err = fmt.Errorf("astiffmpeg: adapting cmd for input #%d failed: %w", idx, err)
			return
		}
	}

	// Output
	if err = out.adaptCmd(cmd); err != nil {
		err = fmt.Errorf("astiffmpeg: adapting cmd for output failed: %w", err)
		return
	}

	// Run cmd
	if err = cmd.Run(); err != nil {
		err = fmt.Errorf("astiffmpeg: running %s failed with stderr %s: %w", strings.Join(cmd.Args, " "), bufErr.Bytes(), err)
		return
	}
	return
}
