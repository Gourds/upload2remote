package providers

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/wonderivan/logger"
	"path"
)

func (c OSS) Auth() (SessionType, error) {
	logger.Debug("Alibaba login ...")
	client, err := oss.New(c.Endpoint, c.AccessKeyID, c.AccessKeySecret)
	return SessionType{oss: client},err
}

func (c OSS) UploadFile(objName string, filePath string, client SessionType) (err error) {
	//获取存储空间
	bucket, err := client.oss.Bucket(c.Bucket)
	if err != nil {
		logger.Error(err)
		return
	}
	//storageType := oss.ObjectStorageClass(oss.StorageStandard)
	//objectAcl := oss.ObjectACL(oss.ACLPublicRead)
	err = bucket.PutObjectFromFile(objName, filePath)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Upload %s To %s://%s Success", filePath, c.Type, path.Join(c.Bucket, objName))
	return
}
