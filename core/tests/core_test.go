package tests

import (
	"fmt"
	"log"
	"testing"
	"yt/core"
)

func TestVideo(t *testing.T) {

	opSys, err := core.ExtractOS()
	if err != nil {
		log.Fatal("err extracting OS: ", err)
	}
	outdir, err := core.GetDefaultDownloadDir(opSys)
	if err != nil {
		log.Fatal("err getting default dir: ", err)
	}
	c := true
	VideoOpts := &core.VideoOptions{
		URL:        "https://youtu.be/XHH453A75n4?si=aWKLiXb5eP8A8hbv",
		Quality:    "370",
		Chapters:   &c,
		OutputName: "subTestAZZ",
		OutDir:     outdir,
		OS:         opSys,
	}
	cmd := core.BuildCommand(VideoOpts)
	t.Log("Command is: ", cmd)
	if err := cmd.Run(); err != nil {
		fmt.Println("👺😡☠️ error running command: ", err)
	}

}

func TestAudio(t *testing.T) {
	opSys, err := core.ExtractOS()
	if err != nil {
		log.Fatal("err extracting OS: ", err)
	}
	outdir, err := core.GetDefaultDownloadDir(opSys)
	if err != nil {
		log.Fatal("err getting default dir: ", err)
	}

	AudioOpts := &core.AudioOptions{
		URL:          "https://youtu.be/2j3x0VYnehg?si=GzCh0tUe0pCj9xdw",
		OutputName:   "ozuna10",
		OutDir:       outdir,
		OS:           opSys,
		AudioQuality: 10,
	}
	cmd := core.BuildCommand(AudioOpts)
	t.Log("Command is: ", cmd)
	if err := cmd.Run(); err != nil {
		fmt.Println("👺😡☠️ error running command: ", err)
	}
}

func TestPlaylist(t *testing.T) {
	opSys, err := core.ExtractOS()
	if err != nil {
		log.Fatal("err extracting OS: ", err)
	}
	outdir, err := core.GetDefaultDownloadDir(opSys)
	if err != nil {
		log.Fatal("err getting default dir: ", err)
	}

	playlistOpts := &core.PlaylistOptions{
		IsAudio:            true,
		URL:                "https://www.youtube.com/playlist?list=PL4cUxeGkcC9jTpR5D2z-xy7YRnWh9xnFM",
		OutDir:             outdir,
		PlaylistFolderName: "Net Ninja Platformer", // optional, can be empty
		OS:                 opSys,
		AudioOpts:          &core.AudioOptions{},
	}

	t.Log("the command to run is \n", playlistOpts.BuildCommand())
	cmd := core.BuildCommand(playlistOpts)
	if err := cmd.Run(); err != nil {
		fmt.Println("👺😡☠️ error running command: ", err)
	}

}
