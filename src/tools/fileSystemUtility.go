package tools

import (
	"os"
	"regexp"
)

func isValidDirName(path string) bool {
	extreg := regexp.MustCompile(`(\.[^.]+)$`)
	return !extreg.MatchString(path)
}

func isValidWoolFileName(path string) bool {
	extreg := regexp.MustCompile(`(\.wool)$`)
	return extreg.MatchString(path)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	if err != nil {
		return false
	}
	return true
}

func isDir(path string) bool {
	if !pathExists(path) {
		return false
	}

	info, _ := os.Stat(path)

	return info.IsDir()
}

func isFile(path string) bool {
	if !pathExists(path) {
		return false
	}

	info, _ := os.Stat(path)

	return !info.IsDir()
}
