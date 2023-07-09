package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/Artemides/go-fiber-api/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var MYSQL *gorm.DB

type Config struct {
	DBHost         string `mapstructure:"MYSQL_ROOT_HOST"`
	DBUserName     string `mapstructure:"MYSQL_USER"`
	DBUserPassword string `mapstructure:"MYSQL_ROOT_PASSWORD"`
	DBName         string `mapstructure:"MYSQL_DB"`
	DBPort         string `mapstructure:"MYSQL_PORT"`
	ClientOrigin   string `mapstructure:"CLIENT_ORIGIN"`
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

func ConnectMySQLDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)

	fmt.Println("dsn ", dsn)
	MYSQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect MYSQL\n", err.Error())
		os.Exit(1)
	}

	MYSQL.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	MYSQL.AutoMigrate(&models.Note{})
	log.Println("🚀 Connection Successfully DB")
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	fmt.Println("user", config.DBUserPassword)

	return
}
