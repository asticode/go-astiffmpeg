Use your FFMpeg binary to manipulate your video files

# Usage

WARNING: the code below doesn't handle errors for readibility purposes. However you SHOULD!

```go
// Build astiffmpeg
var f = astiffmpeg.New(astiffmpeg.Configuration{BinaryPath: <your binary path>})

// Make sure stderr is parsed to retrieve ffmpeg progression
f.SetStdErrParser(astiffmpeg.DefaultStdErrParser(time.Second, func(r astiffmpeg.DefaultStdErrResults) {
    astilog.Debugf("time: %s", r.Time.String())
}))
```