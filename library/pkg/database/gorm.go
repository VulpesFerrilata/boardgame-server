package database

import (
	"strings"

	"github.com/VulpesFerrilata/boardgame-server/library/config"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewGorm(config *config.Config) (*gorm.DB, error) {
	return gorm.Open(strings.ToLower(config.SqlSettings.DriverName), config.SqlSettings.DataSource)
}
