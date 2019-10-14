package main

import (
	"fmt"
	"github.com/tetsuzawa/go-3daudio/app/controllers"
	"github.com/tetsuzawa/go-3daudio/app/models"
	"github.com/tetsuzawa/go-3daudio/config"
	"github.com/tetsuzawa/go-3daudio/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(config.Config.MockString)
	fmt.Println(models.DbConnection)

	//insert data to db for examination
	//t := time.Now()
	//entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	//id := ulid.MustNew(ulid.Now(), entropy)
	//hrtf := models.NewHRTF(id.String(), "tetsu", 20, 20, 0, 0.35555)
	//if err :=hrtf.Create(); err != nil{
	//	log.Fatalln(err)
	//}

	//get data for examination
	//hrtf, err := models.GetHRTF("01DQ44KFF4D44TFZA9963GD1VS")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println("hrtf: ", hrtf)

	err := controllers.StartWebServer()
	if err != nil {
		log.Fatalln(err)
	}
}
