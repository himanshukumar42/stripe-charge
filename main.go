package main

import (
	"fmt"

	"github.com/himanshuk42/stripe-charge/src/models"
	"github.com/himanshuk42/stripe-charge/src/router"

	Config "github.com/himanshuk42/stripe-charge/src/config"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("error", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&models.Charge{})
	r := router.ChargeRouter()
	r.Run(":8090")

}
