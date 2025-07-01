package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/E-nkv/vidder/core"
)

// PromptAudioOptions interactively fills AudioOptions for a video URL or playlist audio.
func PromptAudioOptions(url string, outdir string, opSys core.OS, isPlaylist bool) *core.AudioOptions {

	outputName := ""
	if !isPlaylist {
		outputName = PromptString("File name ? (Leave empty to use default of VIDEO'S TITLE) ")
		if strings.TrimSpace(outputName) == "" {
			outputName = "%(title)s"
		}
	}

	opts := []LabelValue{
		{Label: "MP3", Value: "mp3"},
		{Label: "MP4", Value: "mp4"},
		{Label: "AAC", Value: "aac"},
	}
	ext, err := PromptSelectWithValues("Audio extension ?", opts)
	if err != nil {
		fmt.Println("Download canceled 🥲")
		os.Exit(0)
	}

	quality := PromptIntRange("Audio quality (0 best - 10 worst)? ", 0, 10)

	return &core.AudioOptions{
		URL:          url,
		OutputName:   outputName,
		AudioExt:     ext,
		AudioQuality: quality,
		OutDir:       outdir,
		OS:           opSys,
	}
}
