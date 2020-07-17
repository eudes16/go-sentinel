package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

// Connector ...
var Connector *gorm.DB
var err error

// ConnectorConfig represents db configuration
type ConnectorConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

// BuildConnectorConfig ...
func buildConnectorConfig() *ConnectorConfig {
	_ = godotenv.Load(".env")
	connectorConfig := ConnectorConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_SECRET"),
		DBName:   os.Getenv("DB_SCHEMA"),
	}
	return &connectorConfig
}

// ConnectorURL ...
func connectorURL(connectorConfig *ConnectorConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		connectorConfig.Host,
		connectorConfig.Port,
		connectorConfig.User,
		connectorConfig.DBName,
		connectorConfig.Password,
	)
}

func OpenConnection() {
	Connector, err = gorm.Open("postgres", connectorURL(buildConnectorConfig()))

	if err != nil {
		fmt.Println(err)
	}

	Connector.LogMode(true)

}
