package main

import (
	"t-rex_wrapper/front"
	"t-rex_wrapper/logger"
	"t-rex_wrapper/server"
	"t-rex_wrapper/utils"
)

func main() {
	cli := utils.Cli()
	utils.LoadYamlConfig(cli.FilePathConfig)
	logger.InitLogger(utils.Cfg.ApiLogFile)
	if utils.Cfg.FrontEnd == true {
		front.InitFrontEnd()
	}
	server.GoGinServer()
}
