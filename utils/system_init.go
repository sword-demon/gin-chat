package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.SetConfigFile("config/app.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("config init app success")
}

func InitMySQL() {
	dns := viper.GetString("mysql.dns")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.Exec("SELECT 1").Error
	if err != nil {
		panic(err)
	}

	DB = db
}
