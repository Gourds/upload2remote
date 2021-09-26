package main

import (
	"github.com/gourds/upload2remote/config"
	"github.com/gourds/upload2remote/providers"
	"github.com/gourds/upload2remote/util"
	"github.com/wonderivan/logger"
	"path"
)

func main() {
	var data providers.ObManager
	data = providers.GetConfig()
	logger.Info(data)
	auth, err := data.Auth()
	if err != nil {
		logger.Error("登录信息有误")
		return
	}


	files, err := util.ListDir(config.CommonCfg.SrcPath)
	if err != nil {
		logger.Error("源路径错误", err)
	}
	for _, eachFile := range files {
		objPath := path.Join(config.CommonCfg.RemoteRootPath, eachFile)
		data.UploadFile(objPath, eachFile, auth)
	}

}
