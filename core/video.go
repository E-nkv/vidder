package core

import (
	"fmt"
	"path/filepath"
)

/*
VideoOptions Default Values:
Quality      "best"
Subtitles    nil
DownloadSubs  false
Thumbnails  false
Chapters     true
FileType     "mkv"
OutputName "%(title)s @ %(uploader)s" (for example, "Postgresql Complete Tutorial @ Primeagen")   love ya prime
OutDir "./"
URL "NO DEFAULT. REQUIRED"
*/

type VideoOptions struct {
	CommandBuilder
	Quality    string
	Chapters   *bool
	FileType   string
	OutputName string
	OutDir     string
	URL        string
	OS         OS
}

func (opts *VideoOptions) Clone() CommandBuilder {
	cpy := *opts
	return &cpy
}

func (opts *VideoOptions) GetOS() OS         { return opts.OS }
func (opts *VideoOptions) SetURL(url string) { opts.URL = url }
func (opts *VideoOptions) SetDefaults() {
	if opts.Quality == "" {
		opts.Quality = "best"
	}
	if opts.FileType == "" {
		opts.FileType = "mkv"
	}
	if opts.OutputName == "" {
		opts.OutputName = "%(title)s" //maybe add in future -->    '@ %(uploader)s'
	}
	if opts.OutDir == "" {
		panic("OutDir must be specified")
	}

}
func (opts *VideoOptions) BuildCommand() string {

	opts.SetDefaults()

	args := &CommandArgs{"yt-dlp"}

	// Format selection (video quality)
	if opts.Quality != "best" {
		format := fmt.Sprintf("bestvideo[height<=%s]+bestaudio/best", opts.Quality)
		args.Add("-f", format, true)
	}

	// Chapters
	if opts.Chapters != nil && *opts.Chapters {
		args.Add("--embed-chapters")
	}

	// Filetype
	args.Add("-t", opts.FileType)

	// Output template
	filePath := filepath.Join(opts.OutDir, opts.OutputName)
	output := filePath + ".%(ext)s"
	args.Add("-o", output, true)

	// URL
	args.Add(opts.URL)

	return args.Join(" ")
}
