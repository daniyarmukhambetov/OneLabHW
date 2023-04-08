package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hw1/internal/config"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DBCfg.DbHost,
		cfg.DBCfg.DbUser,
		cfg.DBCfg.DbPassword,
		cfg.DBCfg.DbName,
		cfg.DBCfg.DbPort,
		cfg.DBCfg.SSL,
		cfg.Timezone,
	)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
