package util

import (
	"github.com/wonderivan/logger"
	"os"
	"path/filepath"
)

func ListDir(dirPath string) (files []string, err error) {
	_, err = os.Stat(dirPath)
	if err != nil {
		return
	}
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		logger.Error(err)
		return
	}
	return
}
