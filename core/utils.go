package core

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

func joinSubLangs(langs []string) string {
	return strings.Join(langs, ",")
}

func stringStartsWith(str string, subStr string) bool {
	cuttedPart := strings.TrimPrefix(str, subStr)
	return cuttedPart != str

}

type CommandArgs []string

func (args *CommandArgs) Join(separator string) string {
	return strings.Join(([]string)(*args), separator)
}

// Add takes in the key: string, val?:string, and wrapInQuotes?:bool. It does what you would expect it to do. Yes, it would be great to have optional params in go funcs
func (args *CommandArgs) Add(key string, valAndWrap ...any) {
	shouldWrap := false
	val := ""
	if len(valAndWrap) >= 1 {
		actualVal, ok := valAndWrap[0].(string)
		if ok {
			val = actualVal
		}
		if len(valAndWrap) >= 2 {
			actualVal, ok := valAndWrap[1].(bool)
			if ok {
				shouldWrap = actualVal
			}
		}
	}
	if shouldWrap {
		val = fmt.Sprintf(`"%s"`, val)
	}
	*args = append(*args, key, val)

}

func ExtractOS() (OS, error) {
	osStr := runtime.GOOS
	switch osStr {
	case "linux":
		return LINUX, nil
	case "windows":
		return WINDOWS, nil
	case "darwin":
		return MAC, nil
	default:
		return 255, fmt.Errorf("unsupported OS")
	}

}
func GetDefaultDownloadDir(opSys OS) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	var baseDir string
	switch opSys {
	case WINDOWS:
		baseDir = filepath.Join(usr.HomeDir, "Downloads", "vidder")
	case LINUX, MAC:
		baseDir = filepath.Join(usr.HomeDir, "downloads", "vidder")
	default:
		return "", fmt.Errorf("unsupported OS was sent. Dev, fix this")
	}

	//0755 = rwxr-xr-x.
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return "", err
	}

	return baseDir, nil
}
