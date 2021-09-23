package providers

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gourds/upload2remote/config"
	"github.com/wonderivan/logger"
	"net/http"
	"os"
)

func Upload2s3(objname string,filepath string)  {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(config.CommonCfg.Region),
		Credentials: credentials.NewStaticCredentials(
			config.CommonCfg.AccessKeyID,
			config.CommonCfg.AccessKeySecret,
			""),
	})
	if err != nil {
		logger.Fatal(err)
	}

	file, err2 :=  os.Open(filepath)
	if err2 != nil {
		logger.Error(err2)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)
	//uplaod
	_, s3err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:			aws.String(config.CommonCfg.Bucket),
		Key: 			aws.String(objname),
		ACL:			aws.String("private"),
		Body:			bytes.NewReader(buffer),
		ContentLength:  aws.Int64(size),
		ContentType:    aws.String(http.DetectContentType(buffer)),
		ContentDisposition: aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass: aws.String("INTELLIGENT_TIERING"),
	})
	if s3err != nil {
		logger.Error(s3err)
	}
}