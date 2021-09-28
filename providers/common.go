package providers

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gourds/upload2remote/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/ks3sdklib/aws-sdk-go/service/s3"
	"github.com/wonderivan/logger"
)

type ObManager interface {
	Auth() (SessionType, error)
	UploadFile(string, string, SessionType, *config.Multi) error
}

type SessionType struct {
	s3  *session.Session
	ks3 *s3.S3
	obs *obs.ObsClient
	oss *oss.Client
}

type S3 struct {
	config.Config
}

type KS3 struct {
	config.Config
}

type OSS struct {
	config.Config
}

type OBS struct {
	config.Config
}

func GetConfig() (data ObManager){
	switch config.CommonCfg.Type {
	case "OSS":
		data = &OSS{config.CommonCfg}
		logger.Info("阿里云",data)
	case "OBS":
		data = &OBS{config.CommonCfg}
		logger.Info("华为云",data)
	case "S3":
		data = &S3{config.CommonCfg}
		logger.Info("亚马逊",data)
	case "KS3":
		data = &KS3{config.CommonCfg}
		logger.Info("金山云",data)
	}
	return
}
