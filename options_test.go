package astiffmpeg

import (
	"os/exec"
	"testing"

	"github.com/asticode/go-astitools/ptr"
	"github.com/stretchr/testify/assert"
)

func TestGlobalOptions(t *testing.T) {
	o := GlobalOptions{}
	cmd := &exec.Cmd{}
	o.adaptCmd(cmd)
	assert.Equal(t, []string(nil), cmd.Args)
	assert.Equal(t, []string(nil), cmd.Env)
	o.HideBanner = true
	cmd = &exec.Cmd{}
	o.adaptCmd(cmd)
	assert.Equal(t, []string{"-hide_banner"}, cmd.Args)
	assert.Equal(t, []string(nil), cmd.Env)
	o.Report = true
	cmd = &exec.Cmd{}
	o.adaptCmd(cmd)
	assert.Equal(t, []string{"-hide_banner", "-report"}, cmd.Args)
	assert.Equal(t, []string(nil), cmd.Env)
	o.Log = &LogOptions{Level: LogLevelFatal}
	cmd = &exec.Cmd{}
	o.adaptCmd(cmd)
	assert.Equal(t, []string{"-hide_banner", "-loglevel", "fatal", "-report"}, cmd.Args)
	assert.Equal(t, []string(nil), cmd.Env)
	o.Log.Color = astiptr.Bool(false)
	cmd = &exec.Cmd{}
	o.adaptCmd(cmd)
	assert.Equal(t, []string{"-hide_banner", "-loglevel", "fatal", "-report"}, cmd.Args)
	assert.Equal(t, []string{"AV_LOG_FORCE_NOCOLOR=1"}, cmd.Env)
	o.Log.Repeated = true
	cmd = &exec.Cmd{}
	o.adaptCmd(cmd)
	assert.Equal(t, []string{"-hide_banner", "-loglevel", "repeat+fatal", "-report"}, cmd.Args)
	assert.Equal(t, []string{"AV_LOG_FORCE_NOCOLOR=1"}, cmd.Env)
}

func TestNumber(t *testing.T) {
	n := Number{Value: 3}
	assert.Equal(t, "3", n.String())
	n.Prefix = "M"
	assert.Equal(t, "3M", n.String())
	n.BinaryMultiple = true
	assert.Equal(t, "3Mi", n.String())
	n.ByteMultiple = true
	assert.Equal(t, "3MiB", n.String())
}

func TestStreamSpecifier(t *testing.T) {
	n := StreamSpecifier{Type: StreamSpecifierTypeSubtitle}
	assert.Equal(t, "s", n.String())
	n.Index = astiptr.Int(2)
	assert.Equal(t, "s:2", n.String())
	n.Type = ""
	assert.Equal(t, "2", n.String())
}

func TestEncodingOptions(t *testing.T) {
	e := EncodingOptions{}
	cmd := &exec.Cmd{}
	e.adaptCmd(cmd)
	assert.Equal(t, []string(nil), cmd.Args)
	e = EncodingOptions{Codec: &StreamOption{Stream: &StreamSpecifier{Type: StreamSpecifierTypeVideo}, Value: "h264"}}
	cmd = &exec.Cmd{}
	e.adaptCmd(cmd)
	assert.Equal(t, []string{"-codec:v", "h264"}, cmd.Args)
	e = EncodingOptions{Coder: &StreamOption{Stream: &StreamSpecifier{Type: StreamSpecifierTypeVideo}, Value: true}}
	cmd = &exec.Cmd{}
	e.adaptCmd(cmd)
	assert.Equal(t, []string{"-coder:v", "1"}, cmd.Args)
	e = EncodingOptions{CRF: astiptr.Int(1)}
	cmd = &exec.Cmd{}
	e.adaptCmd(cmd)
	assert.Equal(t, []string{"-crf", "1"}, cmd.Args)
	e = EncodingOptions{Level: astiptr.Float(3)}
	cmd = &exec.Cmd{}
	e.adaptCmd(cmd)
	assert.Equal(t, []string{"-level", "3.0"}, cmd.Args)
	e = EncodingOptions{Preset: &StreamOption{Stream: &StreamSpecifier{Type: StreamSpecifierTypeVideo}, Value: PresetFast}}
	cmd = &exec.Cmd{}
	e.adaptCmd(cmd)
	assert.Equal(t, []string{"-preset:v", "fast"}, cmd.Args)
	e = EncodingOptions{Profile: &StreamOption{Stream: &StreamSpecifier{Type: StreamSpecifierTypeVideo}, Value: ProfileBaseline}}
	cmd = &exec.Cmd{}
	e.adaptCmd(cmd)
	assert.Equal(t, []string{"-profile:v", "baseline"}, cmd.Args)
	e = EncodingOptions{Tune: TuneAnimation}
	cmd = &exec.Cmd{}
	e.adaptCmd(cmd)
	assert.Equal(t, []string{"-tune", "animation"}, cmd.Args)
}
