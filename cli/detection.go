package cli

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"
)

// DetectURLType detects if the URL is a video or playlist.
func DetectURLType(url string) string {
	if strings.Contains(url, "playlist?list=") {
		return "playlist"
	}
	return "video"
}

// GetVideoMetadata calls yt-dlp to get title and uploader for a URL.
func GetVideoMetadata(url string) (string, string, error) {
	titleCmd := exec.Command("yt-dlp", "--get-title", url)
	titleOut, err := titleCmd.Output()
	if err != nil {
		return "", "", err
	}
	uploaderCmd := exec.Command("yt-dlp", "--get-uploader", url)
	uploaderOut, err := uploaderCmd.Output()
	if err != nil {
		return "", "", err
	}
	title := strings.TrimSpace(string(titleOut))
	uploader := strings.TrimSpace(string(uploaderOut))
	if title == "" || uploader == "" {
		return "", "", errors.New("could not retrieve metadata")
	}
	return title, uploader, nil
}

// GetVideoResolutions returns available video resolutions by parsing yt-dlp -F output.
func GetVideoResolutions(url string) ([]string, error) {
	cmd := exec.Command("yt-dlp", "-F", url)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(out), "\n")
	var resolutions []string
	re := regexp.MustCompile(`\d+\s+video.*?(\d{3,4})p`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) == 2 {
			height := matches[1]
			if !contains(resolutions, height) {
				resolutions = append(resolutions, height)
			}
		}
	}
	return resolutions, nil
}

func contains(slice []string, val string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}
