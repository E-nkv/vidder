package cli

import (
	"fmt"
	"os"
	"strings"
	"yt/core"
)

// PromptVideoOptions interactively fills VideoOptions for a video URL.
func PromptVideoOptions(url, outdir string, opSys core.OS, isPlaylist bool) *core.VideoOptions {
	opts := []LabelValue{
		{Label: "Best 🚀", Value: "best"},
		{Label: "1080p 😉", Value: "1080"},
		{Label: "720p 🫡", Value: "720"},
		{Label: "480p 🙂", Value: "480"},
		{Label: "360p 🫤", Value: "360"},
		{Label: "240p 🤮", Value: "240"},
	}

	selectedRes, err := PromptSelectWithValues("Preferred video resolution? ", opts)
	if err != nil {
		fmt.Println("Download canceled 🥲")
		os.Exit(0)
	}

	chapOpts := []LabelValue{
		{Label: "Yep 😎", Value: "yes"},
		{Label: "Nop 🥸", Value: "no"},
	}
	chaps, err := PromptSelectWithValues("Include chapters?", chapOpts)
	if err != nil {
		fmt.Println("Download canceled 🥲")
		os.Exit(0)
	}

	outputName := ""
	if !isPlaylist {
		outputName = PromptString("Video name? (Leave empty to use default name of VIDEO'S TITLE): ")
		outputName = strings.TrimSpace(outputName)

	}
	ext, err := PromptSelectWithValues("Select file extension:", []LabelValue{
		{Label: "MKV (supports chapters / timestamps)", Value: "mkv"},
		{Label: "MP4 (works on all devices)", Value: "mp4"},
	})
	if err != nil {
		fmt.Println("Download canceled 🥲")
		os.Exit(0)
	}
	includeChapters := true
	if chaps == "no" {
		includeChapters = false
	}

	return &core.VideoOptions{
		URL:        url,
		Quality:    selectedRes,
		Chapters:   &includeChapters,
		OutputName: outputName,
		FileType:   ext,
		OutDir:     outdir,
		OS:         opSys,
	}
}
