package main

import (
	"context"
	"github.com/joho/godotenv"
	"hw1/config"
	handler2 "hw1/internal/handler"
	"hw1/internal/service"
	"hw1/internal/storage"
	"hw1/internal/storage/postgres"
	"hw1/internal/transport"
	"log"
	"os"
	"os/signal"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	cfg := config.Config{
		DBCfg: config.DBConfig{
			DbName:     config.GetEnv("DB_NAME", "onelab"),
			DbHost:     config.GetEnv("DB_HOST", "localhost"),
			DbUser:     config.GetEnv("DB_USER", "admin"),
			DbPassword: config.GetEnv("DB_PASSWORD", "password"),
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
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	go func() {
		log.Println(<-sig)
		cancel()
	}()
	err = srv.Run()
	if err != nil {
		log.Fatalln(err)
	}
	<-srv.Conn
}
