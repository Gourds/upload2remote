package providers

import (
	"bytes"
	"github.com/ks3sdklib/aws-sdk-go/aws"
	"github.com/ks3sdklib/aws-sdk-go/aws/credentials"
	"github.com/ks3sdklib/aws-sdk-go/service/s3"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"os"
	"path"
)

func (c KS3) Auth() (SessionType, error) {
	credentials := credentials.NewStaticCredentials(c.AccessKeyID, c.AccessKeySecret, "")
	client := s3.New(&aws.Config{
		Region:           c.Region,
		Credentials:      credentials,
		Endpoint:         c.Endpoint,
		DisableSSL:       true,
		LogLevel:         0,
		S3ForcePathStyle: false,
		LogHTTPBody:      true,
		Logger:           os.Stdout,
	})
	return SessionType{ks3: client}, nil
}

func (c KS3) UploadFile(objName string, filePath string, client SessionType) (err error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Error(err)
		return
	}
	params := &s3.PutObjectInput{
		Bucket:      aws.String(c.Bucket),
		Key:         aws.String(objName),
		ACL:         aws.String("private"),
		Body:        bytes.NewReader(file),
		ContentType: aws.String("application/octet-stream"),
		Metadata:    map[string]*string{
			//"key" : aws.String("metadataValue"), //元数据
		},
	}
	_, err = client.ks3.PutObject(params)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Upload %s To %s://%s Success", filePath, c.Type, path.Join(c.Bucket, objName))
	return
}
