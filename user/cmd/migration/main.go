package main

import (
	"log"

	"github.com/VulpesFerrilata/boardgame-server/library/config"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/database"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/model"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewGorm(cfg)
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
