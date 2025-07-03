package cli

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"sync"

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
		fmt.Println(`Expected usage is: "vidder <URL1> <URL2> <URL3> OR vidder -f path/to/file.txt"`)
		phonkSong := "vidder https://youtu.be/LIddkGIDwJ4?si=UdelqvTvuzgGzAZ0"
		fmt.Printf("Want an example? Run the following command in the terminal:\n\x1b[35m%s\x1b[0m\n\n", phonkSong)
		fmt.Printf("Videos will be downloaded at your downloads/vidder, regardless of your Operative System: \n(C:\\Users\\YOURUSERNAME\\Downloads\\vidder) on windows, and /home/YOURUSERNAME/downloads/vidder in UNIX based OS \n\n")
		fmt.Printf("you can check the project on github at https://github.com/E-nkv/vidder\n")
		return
	}
	figure.NewColorFigure("VIDDER-CLI", "", "blue", true).Print()
	println()
	println()

	if len(args) == 1 {
		url := args[0]
		handleOne(url)
		return
	}

	handleFromFile(args)

}

func handleOne(url string) {
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

func handleFromFile(args []string) {
	if len(args) >= 3 || (len(args) == 2 && args[0] != "-f") {
		fmt.Println("👺 Usage: vidder <URL>, or vidder -f <absolute-path/to/file-that-contains-video-urls.txt>")
		return
	}
	errorWithFile := false
	err := fmt.Errorf("dummy error")
	if errorWithFile {
		fmt.Printf("😡 error parsing file that contains urls: %v\n", err)
		return
	}
	pathToFile := args[1]

	urls, err := ReadURLsFromFile(pathToFile)
	if err != nil {
		fmt.Printf("😡 error opening txt file located at '%s'. Are you sure it exists? (it must be absolute-path/to/file/FILENAME.txt). and if so, is it of the valid format? (list of urls separated by end of line)", pathToFile)
		return
	}
	if len(urls) == 0 {
		fmt.Printf("☠️ it seems there are no URLS in the txt. are you sure thats the file you meant?\n")
		return
	}
	fmt.Printf("BATCH DOWNLOAD OF %v URLS FROM FILE LOCATED AT '%s' 🍖\n", len(urls), pathToFile)
	fmt.Printf("⚠️  the options you will now select will apply to ALL of the downloads in the file  ⚠️\n")

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
		opts = PromptVideoOptions("dummyURL", outdir, opSys, false) //dummy since we'll change it manually
	case "audio":
		opts = PromptAudioOptions("dummyURL", outdir, opSys, false)
	}

	//one minute of hate to the float64-int conversion below 😡😡😡
	concurrentWorkers := int(math.Min(float64(runtime.NumCPU()*5), float64(len(urls))))
	tasks := make(chan string, len(urls))
	var wg sync.WaitGroup
	wg.Add(concurrentWorkers)
	for _, url := range urls {
		tasks <- url
	}
	for id := range concurrentWorkers {
		go downloadWorker(id, opts, tasks, &wg)
	}
	close(tasks)
	wg.Wait()

}
func downloadWorker(id int, opts core.CommandBuilder, tasks <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	//the task is the url
	for url := range tasks {
		fmt.Printf("🤖 WORKER %v will handle URL: %s\n", id, url)
		optsCopy := opts.Clone()
		optsCopy.SetURL(url)
		cmd := core.BuildCommand(optsCopy)
		cmd.Stdout = os.Stdout //pretty uggly, though at least this way the user will know that we are downloading his stuff.
		//he could implement some detailed logging such that each cmd writes to a stream, having a service that contiously reads from the stream and sends the aggregated data to os.Stdout (like download %, errors, and so on)
		cmd.Stderr = nil

		if err := cmd.Run(); err != nil {
			fmt.Printf("😡 worker %v error running command for url [%s], %v. 😡\n", id, url, err)
			continue
		}

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
