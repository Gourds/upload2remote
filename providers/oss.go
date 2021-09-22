package providers
import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/wonderivan/logger"
	"github.com/gourds/upload2remote/config"
)

func Upload2oss(objname string,filepath string){
	cfg := config.CommonCfg
	client, err := oss.New(cfg.Endpoint, cfg.AccessKeyID, cfg.AccessKeySecret)
	if err != nil {
			logger.Error(err)
	}
	//获取存储空间
	bucket, err := client.Bucket(cfg.Bucket)
	if err != nil {
		logger.Error(err)
	}
	//storageType := oss.ObjectStorageClass(oss.StorageStandard)
	//objectAcl := oss.ObjectACL(oss.ACLPublicRead)
	err = bucket.PutObjectFromFile(objname, filepath)
	if err != nil {
		logger.Error(err)
	}
}
