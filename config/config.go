package config

import (
	"fmt"
	"github.com/wonderivan/logger"
)

var (
	CommonCfg   Config
	InputConfig Config
)

type Config struct {
	Type            string `toml:"type"`
	Endpoint        string `toml:"endpoint"`
	AccessKeyID     string `toml:"accessKeyID"`
	AccessKeySecret string `toml:"accessKeySecret"`
	SrcPath         string `toml:"SrcPath"`
	DestPath        string `toml:"DestPath"`
	Bucket          string `toml:"Bucket"`
	RemoteRootPath  string `toml:"RemoteRootPath"`
	Region          string `toml:"Region"`
}

func GetConfig() (err error ){
	initConfig()
	if CommonCfg.DestPath == "" || CommonCfg.SrcPath == "" {
		err = fmt.Errorf("没有配置源或目标路径")
		return err
	}
	return
}

func init(){
	err := GetConfig()
	if err != nil{
		logger.Fatal("初始化配置失败", err)
	}
	logger.Info("初始化配置成功")
}