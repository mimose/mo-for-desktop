package storage

import (
	"os"
	"strings"
)

const localPath = "\\AppData\\Local\\mo"

var localDir string
var localGroupDir string
var localGroupRecordRelDir string
var localRecordDir string

func LocalDir() string {
	if localDir != "" {
		return localDir
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	localDir = strings.Join([]string{homeDir, localPath}, "")
	return localDir
}

func LocalGroupDir() string {
	if localGroupDir != "" {
		return localGroupDir
	}
	localGroupDir = strings.Join([]string{LocalDir(), "\\group"}, "")
	return localGroupDir
}

func LocalGroupDirByKey(key string) (string, string) {
	return LocalGroupDir(), "\\" + key
}

func LocalRecordDir() string {
	if localRecordDir != "" {
		return localRecordDir
	}
	localRecordDir = strings.Join([]string{LocalDir(), "\\record"}, "")
	return localRecordDir
}

func LocalRecordDirByKey(key string) string {
	return strings.Join([]string{LocalRecordDir(), key}, "\\")
}

func LocalGroupRecordRelDir() string {
	if localGroupRecordRelDir != "" {
		return localGroupRecordRelDir
	}
	localGroupRecordRelDir = strings.Join([]string{LocalGroupDir(), "\\recordRel"}, "")
	return localGroupRecordRelDir
}
