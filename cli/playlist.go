package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/E-nkv/vidder/core"
)

// PromptPlaylistOptions interactively fills PlaylistOptions for a playlist URL.
func PromptPlaylistOptions(url, outdir string, opSys core.OS) *core.PlaylistOptions {
	opts := []LabelValue{
		{Label: "Video 📽️", Value: "video"},
		{Label: "Audio 🔉", Value: "audio"},
	}
	choice, err := PromptSelectWithValues("Download playlist as?", opts)
	if err != nil {
		fmt.Println("Download canceled 🥲")
		os.Exit(0)
	}

	playlistFolderName := PromptString("Folder name ? (Leave empty to use default of PLAYLIST'S TITLE) ")
	playlistFolderName = strings.TrimSpace(playlistFolderName)
	if choice == "audio" {
		audioOpts := PromptAudioOptions(url, "", opSys, true) // OutDir set by PlaylistOptions
		return &core.PlaylistOptions{
			URL:                url,
			IsAudio:            true,
			OutDir:             outdir,
			PlaylistFolderName: playlistFolderName,
			AudioOpts:          audioOpts,
			OS:                 opSys,
		}
	} else {
		videoOpts := PromptVideoOptions(url, "", opSys, true) // OutDir set by PlaylistOptions
		return &core.PlaylistOptions{
			URL:                url,
			IsAudio:            false,
			OutDir:             outdir,
			PlaylistFolderName: playlistFolderName,
			VideoOpts:          videoOpts,
			OS:                 opSys,
		}
	}
}
