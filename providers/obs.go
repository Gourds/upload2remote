package providers

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/wonderivan/logger"
	"github.com/gourds/upload2remote/config"
)

func Upload2obs(objname string,filepath string)  {
	var obsClient, _ = obs.New(config.CommonCfg.AccessKeyID,
		config.CommonCfg.AccessKeySecret,
		config.CommonCfg.Endpoint)
	input := &obs.PutFileInput{}
	input.Bucket = config.CommonCfg.Bucket
	input.Key = objname
	input.SourceFile = filepath
	_, err := obsClient.PutFile(input)

	if err != nil {
		logger.Error(err)
	}
}