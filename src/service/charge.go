package Service

import (
	_ "github.com/go-sql-driver/mysql"
	Config "github.com/himanshuk42/stripe-charge/src/config"
	"github.com/himanshuk42/stripe-charge/src/models"
)

func SavePayment(charge *models.Charge) (err error) {
	if err = Config.DB.Create(charge).Error; err != nil {
		return err
	}
	return nil

}
