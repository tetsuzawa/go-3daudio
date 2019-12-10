package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"

	"github.com/tetsuzawa/go-3daudio/web-app/app/controllers"
	"github.com/tetsuzawa/go-3daudio/web-app/app/models"
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
	log.Println("create and insert HRTF")
	hrtf := models.NewHRTF("01DQ44KFF4D44TFZA9963GD1VS", "test_hrtf_kosen", "", 20, 0, 0.35555)
	if err := hrtf.Create(); err != nil {
		log.Fatalln(errors.Wrap(err, "failed to create new HRTF instance in main()"))
	}

	//get data for examination
	//log.Println("get HRTF")
	//hrtf, err := models.GetHRTF("01DQ44KFF4D44TFZA9963GD1VS")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println("hrtf: ", hrtf)

	/*
	//sessiontest
	log.Println("create and insert session")
	session := models.NewSession("01DQ44KFF4D44TFZA9963GD1VS", "tetsuzawa", time.Now())
	if err := session.Create(); err != nil {
		log.Fatalln(errors.Wrap(err, "failed to create new session instance in main()"))
	}

	//get data for examination
	log.Println("get session")
	session, err := models.GetSession("01DQ44KFF4D44TFZA9963GD1VS")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("session: ", session)
	 */

	log.Println("start web server")
	err = controllers.StartWebServer()
	if err != nil {
		log.Fatalln(err)
	}
}
