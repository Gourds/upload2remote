package providers

import (
	"github.com/wonderivan/logger"
	"gourds.site/upload2remote/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
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