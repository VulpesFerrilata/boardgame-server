package main

import (
	"log"

	"github.com/VulpesFerrilata/boardgame-server/library/config"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/database/gorm"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/model"
)

func main() {
	dbConfig, err := config.NewDatabaseConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.NewGorm(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	db.DropTableIfExists(
		&model.User{},
	)

	db.CreateTable(
		&model.User{},
	)
}
