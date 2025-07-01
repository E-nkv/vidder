package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/E-nkv/vidder/core"

	"github.com/common-nighthawk/go-figure"
)

func wantsHelp(args []string) bool {
	if len(args) == 0 {
		return true
	}
	arg := args[0]
	return arg == "--help" || arg == "-h"

}

// Run is the main entry point for the CLI.
// It expects the first argument to be the URL.
func Run() {
	args := os.Args[1:]
	if wantsHelp(args) {
		fmt.Printf("WELCOME TO VIDDER CLI 👋!\n\n")
		fmt.Println(`Expected usage is: "vidder <URL>"`)
		phonkSong := "vidder https://youtu.be/LIddkGIDwJ4?si=UdelqvTvuzgGzAZ0"
		fmt.Printf("Want an example? Run the following command in the terminal:\n\x1b[35m%s\x1b[0m\n\n", phonkSong)
		fmt.Printf("Videos will be downloaded at your downloads/vidder, regardless of your Operative System: \n(C:\\Users\\YOURUSERNAME\\Downloads\\vidder) on windows, and /home/YOURUSERNAME/downloads/vidder in UNIX based OS \n\n")
		fmt.Printf("Wanna check the project on github? https://github.com/E-nkv/vidder\n")
		return
	}
	figure.NewColorFigure("VIDDER-CLI", "", "blue", true).Print()
	println()
	println()
	if len(args) > 1 {
		fmt.Println("👺 Usage: vidder <URL>")
		return
	}
	url := os.Args[1]
	fmt.Printf("PROVIDED URL 🍖: <%s>\n", url)
	urlType := DetectURLType(url)
	switch urlType {
	case "video":
		handleVideoURL(url)
	case "playlist":
		handlePlaylistURL(url)
	default:
		log.Fatalf("Unsupported URL type: %s", urlType)
	}
}

func handleVideoURL(url string) {
	videoOpts := []LabelValue{
		{Label: "Video 📽️", Value: "video"},
		{Label: "Audio 🔉", Value: "audio"},
	}
	choice, err := PromptSelectWithValues("Download as video or audio? ", videoOpts)
	if err != nil {
		fmt.Println("Download canceled 🙂")
		os.Exit(0)
	}

	opSys, err := core.ExtractOS()
	if err != nil {
		log.Fatal("Error extracting OS:", err)
	}
	outdir, err := core.GetDefaultDownloadDir(opSys)
	if err != nil {
		log.Fatal("Error getting default download directory:", err)
	}

	var opts core.CommandBuilder
	switch choice {
	case "video":
		opts = PromptVideoOptions(url, outdir, opSys, false)
	case "audio":
		opts = PromptAudioOptions(url, outdir, opSys, false)
	}
	fmt.Printf("\033[34mStarting %s download... \033[0m\n%s\n", choice, opts.BuildCommand())
	cmd := core.BuildCommand(opts)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running command to download %s: %v", choice, err)
	}
	fmt.Printf("\033[34m%s has been downloaded successfully! \033[0m 😉\n", choice)
}

func handlePlaylistURL(url string) {
	opSys, err := core.ExtractOS()
	if err != nil {
		log.Fatal("Error extracting OS:", err)
	}
	outdir, err := core.GetDefaultDownloadDir(opSys)
	if err != nil {
		log.Fatal("Error getting default download directory:", err)
	}

	po := PromptPlaylistOptions(url, outdir, opSys)

	fmt.Printf("\033[34mStarting playlist download... \033[0m\n%s\n", po.BuildCommand())
	cmd := core.BuildCommand(po)
	if err := cmd.Run(); err != nil {
		log.Fatal("😡 Error running command: ", err)
	}
	fmt.Printf("\033[34mplaylist has been downloaded successfully! \033[0m 😉\n")
}
