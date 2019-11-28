package main

import (
	"log"

	"github.com/tetsuzawa/go-3daudio/web-app/app/controllers"
	"github.com/tetsuzawa/go-3daudio/web-app/config"
	"github.com/tetsuzawa/go-3daudio/web-app/utils"
)

func main() {
	utils.LoggingSettings(config.Cfg.Log.LogFile)
	//fmt.Println(config.Cfg.MockString)
	//log.Println(models.DbConnection)

	//insert data to db for examination
	//t := time.Now()
	//entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	//id := ulid.MustNew(ulid.Now(), entropy)
	//hrtf := models.NewHRTF(id.String(), "tetsu", 20, 20, 0, 0.35555)
	//hrtf := models.NewHRTF("01DQ44KFF4D44TFZA9963GD1VS", "tetsu", 20, 20, 0, 0.35555)
	//if err := hrtf.Create(); err != nil {
	//	log.Fatalln(err)
	//}

	//get data for examination
	//hrtf, err := models.GetHRTF("01DQ44KFF4D44TFZA9963GD1VS")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println("hrtf: ", hrtf)

	log.Println("start web server")
	err := controllers.StartWebServer()
	if err != nil {
		log.Fatalln(err)
	}
}
