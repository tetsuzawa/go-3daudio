package main

import (
	"fmt"
	"log"

	"github.com/tetsuzawa/go-3daudio/app/controllers"
	"github.com/tetsuzawa/go-3daudio/app/models"
	"github.com/tetsuzawa/go-3daudio/config"
	"github.com/tetsuzawa/go-3daudio/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(config.Config.MockString)
	fmt.Println(models.DbConnection)
	//hrtf := models.NewHRTF(1, "tetsu", 20, 20, 0, 0.35555)
	//hrtf := models.GetHRTF("1")
	//fmt.Println("ssss", hrtf)

	err := controllers.StartWebServer()
	if err != nil {
		log.Fatalln(err)
	}
}
