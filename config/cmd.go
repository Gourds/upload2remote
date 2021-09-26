package config

import (
	"github.com/BurntSushi/toml"
	"github.com/gourds/upload2remote/version"
	"github.com/urfave/cli"
	"github.com/wonderivan/logger"
	"os"
	"regexp"
)

func cliConfig() {
	app := &cli.App{
		Name:    "upload2remote",
		Usage:   "Upload local file to remote object stroage",
		Version: version.Version,
		Author:  version.Author,
		Action: func(c *cli.Context) error { //该命令的执行动作函数
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "type,t",
				Value:       "",
				Usage:       "support{OSS|OBS|S3|KS3}",
				Destination: &InputConfig.Type,
			},
			cli.StringFlag{
				Name:        "Endpoint",
				Value:       "",
				Usage:       "endpoint address",
				Destination: &InputConfig.Endpoint,
			},
			cli.StringFlag{
				Name:        "AccessKeyID",
				Value:       "",
				Usage:       "AccessKeyID",
				Destination: &InputConfig.AccessKeyID,
			},
			cli.StringFlag{
				Name:        "AccessKeySecret",
				Value:       "",
				Usage:       "AccessKeySecret",
				Destination: &InputConfig.AccessKeySecret,
			},
			cli.StringFlag{
				Name:        "SrcPath,s",
				Value:       "",
				Usage:       "src file path or dirctory",
				Destination: &InputConfig.SrcPath,
				Required:    true,
			},
			cli.StringFlag{
				Name:        "DestPath,d",
				Value:       "",
				Usage:       "dest path",
				Destination: &InputConfig.DestPath,
				Required:    true,
			},
			cli.StringFlag{
				Name:        "Region",
				Value:       "",
				Usage:       "region only s3 need",
				Destination: &InputConfig.Region,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug(InputConfig)
}

func getBucket()(map[string]string){
	//regexp get bucket name from dst path
	reg1 := regexp.MustCompile(`^(?P<rtype>\w+)://(?P<bucket>.*?)/(?P<rpath>.*)$`)
	logger.Debug(reg1)
	rst := reg1.FindStringSubmatch(CommonCfg.DestPath)
	groupNames := reg1.SubexpNames()
	gnd := make(map[string]string)
	for i, name := range groupNames {
		if i !=0 && name != "" {
			gnd[name] = rst[i]
		}
	}
	CommonCfg.Bucket = gnd["bucket"]
	CommonCfg.RemoteRootPath = gnd["rpath"]
	logger.Debug(gnd)
	return gnd
}

func localConfig() {
	if _, err := toml.DecodeFile("conf/config.toml", &CommonCfg); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	logger.Debug(CommonCfg)
}

func initConfig() {
	cliConfig()
	localConfig()
	if InputConfig.SrcPath != ""{
		CommonCfg.SrcPath = InputConfig.SrcPath
	}
	if InputConfig.DestPath != ""{
		CommonCfg.DestPath = InputConfig.DestPath
		getBucket()
	}
	if InputConfig.AccessKeyID != ""{
		CommonCfg.AccessKeyID = InputConfig.AccessKeyID
	}
	if InputConfig.AccessKeySecret != ""{
		CommonCfg.AccessKeySecret = InputConfig.AccessKeySecret
	}
	if InputConfig.Endpoint != ""{
		CommonCfg.Endpoint = InputConfig.Endpoint
	}
	if InputConfig.Type != "" {
		CommonCfg.Type = InputConfig.Type
	}
	if InputConfig.Region != "" {
		CommonCfg.Region = InputConfig.Region
	}
	logger.Info("配置信息：",CommonCfg)
}

