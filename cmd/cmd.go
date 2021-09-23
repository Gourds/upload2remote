package cmd

import (
	"github.com/gourds/upload2remote/config"
	"github.com/urfave/cli"
	"github.com/wonderivan/logger"
	"os"
)

var InputConfig config.Config

var INFO config.Info


func CliConfig() {
	app := &cli.App{
		Name:    "upload2remote",
		Usage:   "Upload local file to remote object stroage",
		Version: INFO.Version,
		Author:  INFO.Auther,
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
