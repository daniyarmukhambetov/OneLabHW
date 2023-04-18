package postgres

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm/utils/tests"
	"hw1/config"
	"hw1/models"
	"hw1/storage/postgres/mock"
	"testing"
)

func TestUser_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	want := []models.User{
		{
			ID:       3,
			Username: "daniyar2",
			Email:    "daniyar20507050711@gmail.com",
			Name:     "daniyar22",
			LastName: "bakhtiyar22",
			Password: "$2a$14$/xuW3MPj6F.z9q3aVPH1vehlmmG7y8oCxNL7P14lmkjh4OxOW5m.a",
		},
		{
			ID:       4,
			Username: "daniyarello",
			Email:    "daniyarello@mail.ru",
			Name:     "dan",
			LastName: "bak",
			Password: "$2a$14$DxfHC7JHFCIk7sUhClJii.8WRuoDvWd51sO5PxpBKI6gj5DsdBM8i",
		},
		{
			ID:       1,
			Username: "daniyar",
			Email:    "daniyar0507050711@gmail.com",
			Name:     "dan",
			LastName: "bak",
			Password: "$2a$14$B4htPNLEpO/9M./XpHDhEOtgEv3fw6N/L/U2cwYfWqJi5LFqmBCo2",
		},
	}
	UserRepo := mock.NewMockIUser(ctrl)
	cfg := config.Config{
		DBCfg: config.DBConfig{
			DbName:     "onelab",
			DbHost:     "localhost",
			DbUser:     "admin",
			DbPassword: "password",
			DbPort:     "5432",
			SSL:        "disable",
		},
		Addr:      "",
		Timezone:  "Asia/Almaty",
		JWTSecret: nil,
	}
	db, err := InitDB(&cfg)
	if err != nil {
		t.Error(err)
		return
	}
	var users []models.User
	err = db.Raw("SELECT * FROM USERS;").Scan(&users).Error
	UserRepo.EXPECT().List().Return(users, err)
	res, err := UserRepo.List()
	fmt.Println(res)
	tests.AssertEqual(t, res, want)
	//tests.AssertEqual()
}
