package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// Connector ...
var Connector *gorm.DB

// ConnectorConfig represents db configuration
type ConnectorConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

// BuildConnectorConfig ...
func BuildConnectorConfig() *ConnectorConfig {
	connectorConfig := ConnectorConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_SECRET"),
		DBName:   os.Getenv("DB_SCHEMA"),
	}
	return &connectorConfig
}

// ConnectorURL ...
func ConnectorURL(connectorConfig *ConnectorConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		connectorConfig.Host,
		connectorConfig.User,
		connectorConfig.Password,
		connectorConfig.Port,
		connectorConfig.DBName,
	)
}
