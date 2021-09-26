package providers

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/wonderivan/logger"
	"net/http"
	"os"
	"path"
)

func (c S3) Auth() (SessionType, error) {
	client, err := session.NewSession(&aws.Config{
		Region: aws.String(c.Region),
		Credentials: credentials.NewStaticCredentials(
			c.AccessKeyID,
			c.AccessKeySecret,
			""),
	})
	return SessionType{s3: client}, err
}

func (c S3) UploadFile(objName string, filePath string, client SessionType) (err error){

	file, err := os.Open(filePath)
	if err != nil {
		logger.Error(err)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)
	//upload
	_, err = s3.New(client.s3).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(c.Bucket),
		Key:                  aws.String(objName),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Upload %s To %s://%s Success", filePath, c.Type, path.Join(c.Bucket, objName))
	return
}
