package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"hw1/config"
	_ "hw1/docs"
	handler2 "hw1/handler"
	"hw1/service"
	"hw1/storage"
	"hw1/storage/postgres"
	"hw1/transport"
	"log"
	"os"
	"os/signal"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      127.0.0.1:8080
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	log.Fatalln(fmt.Sprintf("server err %s"), run().Error())
}

func run() error {
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
		return err

	}
	repo, err := storage.NewStorage(db)
	if err != nil {
		return err
	}
	service, err := service.NewManager(repo, &cfg)
	if err != nil {
		return err
	}
	handler, err := handler2.NewHandler(service)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	srv, err := transport.NewServer(&cfg, ctx, handler)
	if err != nil {
		return err
	}
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	go func() {
		log.Println(<-sig)
		cancel()
	}()
	err = srv.Run()
	if err != nil {
		return err
	}
	<-srv.Conn
	return nil
}
