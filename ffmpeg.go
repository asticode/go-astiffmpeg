package astiffmpeg

import (
	"bytes"
	"context"
	"os/exec"
	"strings"

	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// FFMpeg represents an entity capable of running an FFMpeg binary
// https://ffmpeg.org/ffmpeg.html
type FFMpeg struct {
	binaryPath string
}

// New creates a new FFMpeg
func New(c Configuration) *FFMpeg {
	return &FFMpeg{binaryPath: c.BinaryPath}
}

// Exec executes the binary with the specified options
// ffmpeg [global_options] {[input_file_options] -i input_url} ... {[output_file_options] output_url} ...
func (f *FFMpeg) Exec(ctx context.Context, env []string, g GlobalOptions, in []Input, out []Output) (err error) {
	// Create cmd
	var cmd = exec.CommandContext(ctx, f.binaryPath)
	cmd.Env = env
	var bufOut, bufErr = &bytes.Buffer{}, &bytes.Buffer{}
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr

	// Global options
	g.adaptCmd(cmd)

	// Inputs
	for idx, i := range in {
		if err = i.adaptCmd(cmd); err != nil {
			err = errors.Wrapf(err, "astiffmpeg: adapting cmd for input #%d failed", idx)
			return
		}
	}

	// Outputs
	for idx, o := range out {
		if err = o.adaptCmd(cmd); err != nil {
			err = errors.Wrapf(err, "astiffmpeg: adapting cmd for output #%d failed", idx)
			return
		}
	}

	// Run cmd
	astilog.Debugf("Executing %s", strings.Join(cmd.Args, " "))
	if err = cmd.Run(); err != nil {
		err = errors.Wrapf(err, "astiffmpeg: running %s failed with stderr %s", strings.Join(cmd.Args, " "), bufErr.Bytes())
		return
	}

	// TODO Parse results
	return
}
