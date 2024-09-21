package database

import (
    "fmt"
    "log"
    "stock-api/models" // Import your models package
    "github.com/spf13/viper"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
    // Initialize viper to read the config file
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./config")

    err := viper.ReadInConfig()
    if err != nil {
        log.Fatalf("Error reading config file: %s", err)
    }

    // Retrieve database configuration from config.yaml
    host := viper.GetString("database.host")
    user := viper.GetString("database.user")
    password := viper.GetString("database.password")
    dbname := viper.GetString("database.dbname")
    port := viper.GetInt("database.port")
    sslmode := viper.GetString("database.sslmode")

    // Build the DSN (Data Source Name) from the config values
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, user, password, dbname, port, sslmode)

    // Connect to the `stockdb` database
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to 'stockdb':", err)
    }

    // Migrate the Stock model
    DB.AutoMigrate(&models.Stock{})
}