package main

import (
	"fmt"
	"github.com/tetsuzawa/go-3daudio/app/controllers"
	"github.com/tetsuzawa/go-3daudio/app/models"
	"github.com/tetsuzawa/go-3daudio/config"
	"github.com/tetsuzawa/go-3daudio/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(config.Config.MockString)
	fmt.Println(models.DbConnection)

	controllers.StartWebServer()
}