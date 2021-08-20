package main

import (
	"t-rex_wrapper/front"
	"t-rex_wrapper/server"
	"t-rex_wrapper/utils"
)

func main() {
	cli := utils.Cli()
	utils.InitLogger(cli.LogFile)
	utils.LoadYamlConfig(cli.FilePathConfig)
	if cli.RunFront == true {
		front.InitFrontEnd()
	}
	server.GoGinServer()
}
