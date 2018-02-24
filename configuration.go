package astiffmpeg

import "flag"

// Flags
var (
	BinaryPath = flag.String("ffmpeg-binary-path", "", "the FFMpeg binary path")
)

// Configuration represents the ffmpeg configuration
type Configuration struct {
	BinaryPath string `toml:"binary_path"`
}

// FlagConfig generates a Configuration based on flags
func FlagConfig() Configuration {
	return Configuration{
		BinaryPath: *BinaryPath,
	}
}
