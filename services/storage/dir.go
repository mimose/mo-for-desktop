package storage

import (
	"os"
	"strings"
)

const localPath = "/AppData/Local/mo"

var localDir string
var localSpaceDir string
var localSpaceRecordRelDir string
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

func LocalSpaceDir() string {
	if localSpaceDir != "" {
		return localSpaceDir
	}
	localSpaceDir = strings.Join([]string{LocalDir(), "/space"}, "")
	return localSpaceDir
}

func LocalSpaceDirByKey(key string) (string, string) {
	return LocalSpaceDir(), "/" + key
}

func LocalRecordDir() string {
	if localRecordDir != "" {
		return localRecordDir
	}
	localRecordDir = strings.Join([]string{LocalDir(), "/record"}, "")
	return localRecordDir
}

func LocalRecordDirByKey(key string) string {
	return strings.Join([]string{LocalRecordDir(), key}, "/")
}

func LocalSpaceRecordRelDir() string {
	if localSpaceRecordRelDir != "" {
		return localSpaceRecordRelDir
	}
	localSpaceRecordRelDir = strings.Join([]string{LocalSpaceDir(), "/recordRel"}, "")
	return localSpaceRecordRelDir
}
