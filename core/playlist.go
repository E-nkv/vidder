package core

import (
	"path/filepath"
)

type PlaylistOptions struct {
	CommandBuilder
	IsAudio            bool
	VideoOpts          *VideoOptions
	AudioOpts          *AudioOptions
	URL                string
	OutDir             string // base output directory
	PlaylistFolderName string // custom folder name; if empty, use %(playlist_title)s
	OS                 OS
}

func (opts *PlaylistOptions) Clone() CommandBuilder {
	cpy := *opts
	return &cpy
}

func (opts *PlaylistOptions) SetURL(url string) { opts.URL = url }
func (opts *PlaylistOptions) GetOS() OS {
	return opts.OS
}

func (opts *PlaylistOptions) SetDefaults() {
	if opts.OutDir == "" {
		panic("outdir must be specified in playlist")
	}

	// Use custom playlist folder name or default yt-dlp placeholder
	if opts.PlaylistFolderName == "" {
		opts.PlaylistFolderName = "%(playlist_title)s"
	}

	playlistDir := filepath.Join(opts.OutDir, opts.PlaylistFolderName)

	// Initialize and set defaults for contained options
	if opts.IsAudio {
		if opts.AudioOpts == nil {
			panic("audioOpts cannot be nil if IsAudio == true")
		}

		opts.AudioOpts.OutDir = playlistDir
		opts.AudioOpts.OutputName = "%(playlist_index)s. %(title)s"

		opts.AudioOpts.URL = opts.URL
		opts.AudioOpts.OS = opts.OS
	} else {
		if opts.VideoOpts == nil {
			panic("videoOpts cannot be nil if IsAudio is set to false. inside PlaylistOptions")
		}

		opts.VideoOpts.OutDir = playlistDir
		opts.VideoOpts.OutputName = "%(playlist_index)s. %(title)s"
		opts.VideoOpts.URL = opts.URL
		opts.VideoOpts.OS = opts.OS

	}
}

func (opts *PlaylistOptions) BuildCommand() string {
	opts.SetDefaults()

	if opts.IsAudio && opts.AudioOpts != nil {
		return opts.AudioOpts.BuildCommand()
	} else if !opts.IsAudio && opts.VideoOpts != nil {
		return opts.VideoOpts.BuildCommand()
	}

	panic("DEVELOPER ERROR: something wrong in buildCommand for playlist")
}
