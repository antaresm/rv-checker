package adminPanel

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"rv-check/dataProvider"
	"rv-check/models"
	"github.com/qor/admin"
)

func InitAdmin() *http.ServeMux {
	DB := dataProvider.CurrentDb
	DB.AutoMigrate(&models.RvCheck{})

	Admin := admin.New(&admin.AdminConfig{DB: DB})
	Admin.AddResource(&models.RvCheck{})

	m := http.NewServeMux()
	//Admin.MountTo("/admin", m)
	return m
}
