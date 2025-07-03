package core

import (
	"fmt"
	"path/filepath"
)

type AudioOptions struct {
	CommandBuilder
	AudioExt     string // e.g. "mp3", "m4a"
	AudioQuality int    // 0 (best) to 10, default 0
	URL          string
	OutDir       string
	OutputName   string
	OS           OS
}

func (opts *AudioOptions) Clone() CommandBuilder {
	cpy := *opts
	return &cpy
}

func (opts *AudioOptions) SetURL(url string) { opts.URL = url }

func (opts *AudioOptions) GetOS() OS { return opts.OS }
func (opts *AudioOptions) SetDefaults() {
	if opts.AudioExt == "" {
		opts.AudioExt = "mp3"
	}
	if opts.AudioQuality < 0 || opts.AudioQuality > 10 {
		opts.AudioQuality = 0
	}
	if opts.OutputName == "" {
		opts.OutputName = "%(title)s @ %(uploader)s"
	}
	if opts.OutDir == "" {
		panic("OutDir must be specified")
	}
}
func (opts *AudioOptions) BuildCommand() string {
	opts.SetDefaults()

	args := &CommandArgs{"yt-dlp"}

	// Extract audio only
	args.Add("-x")

	// Audio format
	args.Add("-t", opts.AudioExt)

	// Audio quality (0 best, 1-10 lower quality)
	args.Add("--audio-quality", fmt.Sprintf("%d", opts.AudioQuality))

	// Output template with directory + filename + extension
	filePath := filepath.Join(opts.OutDir, opts.OutputName)
	output := filePath + ".%(ext)s"
	args.Add("-o", output, true)

	// URL
	args.Add(opts.URL)

	return args.Join(" ")
}
