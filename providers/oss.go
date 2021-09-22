package providers
import (
	"github.com/wonderivan/logger"
	"gourds.site/upload2remote/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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
