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
	defer util.MeasureTime("main")()
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
	s,f := 0,0
	for _, eachFile := range files {
		s+=1
		objPath := path.Join(config.CommonCfg.RemoteRootPath, eachFile)
		err = data.UploadFile(objPath, eachFile, auth)
		if err != nil {
			f+=1
		}
	}
	logger.Info("总共:%d\t成功:%d\t失败:%d",s,s-f,f)
}
