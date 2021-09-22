package config

import (
	"github.com/BurntSushi/toml"
	"github.com/wonderivan/logger"
	"os"
)

type Info struct {
	Version string
	Auther string
}

type Config struct {
	Type	string	`toml:"type"`
	Endpoint string `toml:"endpoint"`
	AccessKeyID string `toml:"accessKeyID"`
	AccessKeySecret string `toml:"accessKeySecret"`
	SrcPath string `toml:"SrcPath"`
	DestPath string `toml:"DestPath"`
	Bucket string `toml:"Bucket"`
	RemotRootPath string `toml:"RemotRootPath"`
	Region string `toml:"Region"`
}

var (
	CommonCfg Config
)

func LocalConfig() {
	if _, err := toml.DecodeFile("conf/config.toml", &CommonCfg); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	logger.Debug(CommonCfg)
}
