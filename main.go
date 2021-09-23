package main

import (
	"github.com/gourds/upload2remote/cmd"
	"github.com/gourds/upload2remote/config"
	"github.com/gourds/upload2remote/providers"
	"github.com/wonderivan/logger"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

func lastConfig() {
	cmd.CliConfig()
	config.LocalConfig()
	if cmd.InputConfig.SrcPath != ""{
		config.CommonCfg.SrcPath = cmd.InputConfig.SrcPath
	}
	if cmd.InputConfig.DestPath != ""{
		config.CommonCfg.DestPath = cmd.InputConfig.DestPath
	}
	if cmd.InputConfig.AccessKeyID != ""{
		config.CommonCfg.AccessKeyID = cmd.InputConfig.AccessKeyID
	}
	if cmd.InputConfig.AccessKeySecret != ""{
		config.CommonCfg.AccessKeySecret = cmd.InputConfig.AccessKeySecret
	}
	if cmd.InputConfig.Endpoint != ""{
		config.CommonCfg.Endpoint = cmd.InputConfig.Endpoint
	}
	if cmd.InputConfig.Type != "" {
		config.CommonCfg.Type = cmd.InputConfig.Type
	}
	if cmd.InputConfig.Region != "" {
		config.CommonCfg.Region = cmd.InputConfig.Region
	}
	//regexp get bucket name from dst path
	reg1 := regexp.MustCompile(`^(?P<rtype>\w+)://(?P<bucket>.*?)/(?P<rpath>.*)$`)
	if reg1 == nil {
		logger.Error("regexp err")
	}
	rst := reg1.FindStringSubmatch(config.CommonCfg.DestPath)
	groupNames := reg1.SubexpNames()
	gnd := make(map[string]string)
	for i, name := range groupNames {
		if i !=0 && name != "" {
			gnd[name] = rst[i]
		}
	}
	config.CommonCfg.Bucket = gnd["bucket"]
	config.CommonCfg.RemotRootPath = gnd["rpath"]
	logger.Debug(gnd)
	logger.Info(config.CommonCfg)
}

func ListDir(dirPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func UploadFiles(files[]string, upload func(objname string, filepath string)) {
	for _,f := range files{
		var rf1 = path.Join(config.CommonCfg.RemotRootPath, f)
		upload(rf1, f)
		logger.Info("Upload %s to %s://%s/%s Complate", f, config.CommonCfg.Type, config.CommonCfg.Bucket, rf1)
	}
}

func main() {
	lastConfig()
	files,_ := ListDir(config.CommonCfg.SrcPath)
	logger.Info(files)
	switch config.CommonCfg.Type {
	case "OSS":
		logger.Debug("upload to aliyun oss")
		UploadFiles(files, providers.Upload2oss)
	case "OBS":
		logger.Debug("upload to huawei obs")
		UploadFiles(files, providers.Upload2obs)
	case "KS3":
		logger.Debug("upload to jinshan ks3")
		UploadFiles(files, providers.Upload2ks3)
	case "S3":
		logger.Debug("upload to aws s3")
		UploadFiles(files, providers.Upload2s3)
	default:
		logger.Debug("github.com")
	}

}

