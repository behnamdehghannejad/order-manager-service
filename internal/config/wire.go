package config

import (
	"database/sql"
	"log"

	"gorm.io/gorm"
)

func startService() {
	cfg := loadConfiguration()
	initPostgres(cfg)

	//repo := repository.NewTaskRepository(db)
	//service := service2.NewTaskService(repo)
}
func loadConfiguration() *Config {
	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

func initPostgres(cfg *Config) *gorm.DB {
	db, err := NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
