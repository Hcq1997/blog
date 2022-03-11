package config

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetFilePath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}
