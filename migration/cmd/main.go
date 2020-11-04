package main

import (
	"log"

	"github.com/VulpesFerrilata/library/config"
	"github.com/VulpesFerrilata/library/pkg/database"
	"github.com/VulpesFerrilata/migration/model"
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

	if err := db.AutoMigrate(
		&model.User{},
		&model.Claim{},
	); err != nil {
		log.Fatal(err)
	}
}
