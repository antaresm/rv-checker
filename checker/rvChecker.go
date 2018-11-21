package checker

import (
	"time"
	"log"
	"rv-check/config"
	"rv-check/dataProvider"
)

var appCheckers = []AppChecker{}

func Init(){
	go check()
}

func check() {
	for range time.Tick(config.CheckTimer){
		orm := dataProvider.OrmProvider
		checks := orm.GetChecks()
		appCheckers = LoadCheckers(checks)
		cnt := make(chan int)
		for _, u := range appCheckers {
			go u.Check(cnt)
		}
		for i := 0; i < len(appCheckers); i++ {
			<-cnt
		}
	}
}

func CheckClose() {
	log.Println("Close checks")
	for _, u := range appCheckers {
		postStatus(u.Target, 0)
	}
}
