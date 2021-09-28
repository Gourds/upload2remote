package main

import (
	"github.com/gourds/upload2remote/config"
	"github.com/gourds/upload2remote/providers"
	"github.com/gourds/upload2remote/util"
	"github.com/wonderivan/logger"
	"path"
	"sync"
)

func main() {
	var data providers.ObManager
	var wg = config.Multi{WG: sync.WaitGroup{}, UploadResult: [3]int{0, 0, 0}}
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
	for _, eachFile := range files {
		wg.WG.Add(1)
		wg.UploadResult[0] += 1
		objPath := path.Join(config.CommonCfg.RemoteRootPath, eachFile)
		go data.UploadFile(objPath, eachFile, auth, &wg)
	}
	wg.WG.Wait()
	logger.Info("总共:%d\t成功:%d\t失败:%d",wg.UploadResult[0],wg.UploadResult[1],wg.UploadResult[2])
}
