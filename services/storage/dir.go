package storage

import (
	"os"
	"runtime"
	"strings"
)

var localPath string
var localDir string
var localRecordDir string

func LocalDir() string {
	if localDir != "" {
		return localDir
	}
	homeDir, err := os.UserHomeDir()
	if goos := runtime.GOOS; goos != "windows" {
		localPath = "/.mo"
	} else {
		localPath = "/AppData/Local/mo"
	}
	if err != nil {
		panic(err)
	}
	localDir = strings.Join([]string{homeDir, localPath}, "")
	return localDir
}

func LocalRecordDir() string {
	if localRecordDir != "" {
		return localRecordDir
	}
	localRecordDir = strings.Join([]string{LocalDir(), "/record"}, "")
	return localRecordDir
}

func LocalRecordDirByKey(key string) (string, string) {
	return LocalRecordDir(), "/" + key
}
