package dataProvider

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"rv-check/config"
	"rv-check/models"
	_ "github.com/go-sql-driver/mysql"
)

type ormProvider struct {
	db *gorm.DB
}

var OrmProvider ormProvider
var CurrentDb *gorm.DB

func init() {
	time.Sleep(config.StartDelay)

	CurrentDb = getDB()
	OrmProvider = ormProvider{
		db: CurrentDb,
	}
}

func (dbp ormProvider) Close() {
	dbp.db.Close()
}

func getDB() *gorm.DB {
	dbString := config.DbString
	log.Printf("Open DB %s \n", dbString)
	newDb, err := gorm.Open("mysql", dbString)
	if err != nil {
		log.Println(err)
	}
	newDb.DB().SetConnMaxLifetime(time.Second * 5)
	newDb.DB().SetMaxIdleConns(0)
	newDb.DB().SetMaxOpenConns(50)

	if err != nil {
		log.Println(err)
		return nil
	}
	return newDb
}

func (dbp ormProvider) GetChecks() []models.RvCheck {
	checkpoints := []models.RvCheck{}
	dbp.db.Table(models.RvCheck{}.TableName()).Find(&checkpoints)
	return checkpoints
}

func (dbp ormProvider) UpdateCheck(p *models.RvCheck) {
	dbp.db.Save(p)
}