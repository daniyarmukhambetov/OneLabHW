package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"hw1/config"
	handler2 "hw1/internal/handler"
	"hw1/internal/service"
	"hw1/internal/storage"
	"hw1/internal/storage/postgres"
	"hw1/internal/transport"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(config.GetEnv("DB_TYPE", "2"))
	cfg := config.Config{
		DBCfg: config.DBConfig{
			DbName:     config.GetEnv("DB_NAME", "postgres"),
			DbHost:     config.GetEnv("DB_HOST", "postgres"),
			DbUser:     config.GetEnv("DB_USER", "postgres"),
			DbPassword: config.GetEnv("DB_PASSWORD", "postgres"),
			DbPort:     config.GetEnv("DB_PORT", "5432"),
			SSL:        config.GetEnv("SSL", "disable"),
		},
		Addr:      config.GetEnv("ADDR", ":8080"),
		Timezone:  config.GetEnv("TIMEZONE", "Asia/Almaty"),
		JWTSecret: []byte(config.GetEnv("JWT_SECRET", "secret")),
	}
	db, err := postgres.InitDB(&cfg)
	if err != nil {
		log.Fatalln(err)
	}
	repo, err := storage.NewStorage(db)
	if err != nil {
		log.Fatalln(err)
	}
	service, err := service.NewManager(repo, &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	handler, err := handler2.NewHandler(service)
	if err != nil {
		log.Fatalln(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	srv, err := transport.NewServer(&cfg, ctx, handler)
	if err != nil {
		log.Fatalln(err)
	}
	err = srv.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
