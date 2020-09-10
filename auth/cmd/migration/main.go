package main

import (
	"log"

	"github.com/VulpesFerrilata/boardgame-server/auth/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/library/config"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/database"
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
		&model.Token{},
	)

	db.CreateTable(
		&model.Token{},
	)
}
