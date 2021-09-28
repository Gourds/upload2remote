package providers

import (
	"github.com/gourds/upload2remote/config"
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

func (c *OBS) UploadFile(objName string, filePath string, client SessionType, wg *config.Multi) (err error) {
	defer wg.WG.Done()
	input := &obs.PutFileInput{}
	input.Bucket = c.Bucket
	input.Key = objName
	input.SourceFile = filePath
	_, err = client.obs.PutFile(input)

	if err != nil {
		wg.UploadResult[2] += 1
		logger.Error(err)
		return
	}
	wg.UploadResult[1] += 1
	logger.Info("Upload %s To %s://%s Success", filePath, c.Type, path.Join(c.Bucket, objName))
	return
}
