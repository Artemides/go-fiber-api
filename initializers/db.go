package initializers

import (
	"log"
	"os"

	"github.com/Artemides/go-fiber-api/models"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Config struct {
	DBHost         string `mapstructure:"MYSQL_HOST"`
	DBUsername     string `mapstructure:"MSQL_USER"`
	DBUserPassword string `mapstructure:"MYSQL_PASSWORD"`
	DBName         string `mapstructure:"MUSQL_DB"`
	DBPort         string `mapstructure:"MYSQL_PORT"`

	ClientOrigin string `mapstrcuture:"CLIENT_ORIGIN"`
}

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB\n", err.Error())
		os.Exit(1)
	}

	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	DB.AutoMigrate(&models.Note{})
	log.Println("🚀 Connection Successfully DB")

}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
