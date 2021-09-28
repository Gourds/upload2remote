package providers

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/wonderivan/logger"
	"path"
)

func (c *OBS) Auth() (SessionType, error) {
	logger.Debug("HuaWei login ...")
	client, err := obs.New(c.AccessKeyID,
		c.AccessKeySecret,
		c.Endpoint)
	return SessionType{obs: client}, err
}

func (c *OBS) UploadFile(objName string, filePath string, client SessionType) (err error){

	input := &obs.PutFileInput{}
	input.Bucket = c.Bucket
	input.Key = objName
	input.SourceFile = filePath
	_, err = client.obs.PutFile(input)

	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Upload %s To %s://%s Success", filePath, c.Type, path.Join(c.Bucket, objName))
	return
}
