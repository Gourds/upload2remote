package providers

import (
	"bytes"
	"github.com/gourds/upload2remote/config"
	"github.com/ks3sdklib/aws-sdk-go/aws"
	"github.com/ks3sdklib/aws-sdk-go/aws/credentials"
	"github.com/ks3sdklib/aws-sdk-go/service/s3"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"os"
)

func Upload2ks3(objname string,filepath string)  {

	credentials := credentials.NewStaticCredentials(config.CommonCfg.AccessKeyID, config.CommonCfg.AccessKeySecret, "")
	client := s3.New(&aws.Config{
		Region:           config.CommonCfg.Region,
		Credentials:      credentials,
		Endpoint:         config.CommonCfg.Endpoint,
		DisableSSL:       true,
		LogLevel:         0,
		S3ForcePathStyle: false,
		LogHTTPBody:      true,
		Logger:           os.Stdout,
	})

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		logger.Error(err)
	}
	params := &s3.PutObjectInput{
		Bucket: aws.String(config.CommonCfg.Bucket),
		Key: aws.String(objname),
		ACL: aws.String("private"),
		Body: bytes.NewReader(file),
		ContentType: aws.String("application/octet-stream"),
		Metadata: map[string]*string{
			//"key" : aws.String("metadataValue"), //元数据
		},
	}
	_, err = client.PutObject(params)
	if err != nil {
		panic(err)
	}
}