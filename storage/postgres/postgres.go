package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hw1/config"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		cfg.DBCfg.DbUser,
		cfg.DBCfg.DbPassword,
		cfg.DBCfg.DbName,
		cfg.DBCfg.SSL,
		cfg.DBCfg.DbHost,
		cfg.DBCfg.DbPort,
	)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
