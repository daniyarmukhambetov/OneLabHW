package service

import (
	"github.com/golang/mock/gomock"
	"gorm.io/gorm/utils/tests"
	"hw1/models"
	"hw1/service/mock"
	"testing"
)

func TestUserService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userMock := mock.NewMockIUserService(ctrl)
	ret := make([]models.User, 0)
	ret = append(ret, models.User{
		ID:       0,
		Username: "1",
		Email:    "2",
		Name:     "3",
		LastName: "4",
		Password: "5",
	})
	userMock.EXPECT().List().Return(ret, nil)
	res, _ := userMock.List()
	tests.AssertEqual(t, res, ret)
}

func TestUserService_Create(t *testing.T) {
	Tests := []struct {
		name string
		test models.UserModelIn
		want models.User
		err  interface{}
	}{
		{
			name: "User Success Created",
			test: models.UserModelIn{
				Name:     "daniyar",
				LastName: "bakhtiyar",
				Email:    "dan@mail.ru",
				Username: "daniyarello",
				Password: "daniyar",
			},
			want: models.User{
				Name:     "daniyar",
				LastName: "bakhtiyar",
				Email:    "dan@mail.ru",
				Username: "daniyarello",
				Password: "daniyar",
				ID:       1,
			},
			err: nil,
		},
	}

	for _, test := range Tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			UserMock := mock.NewMockIUserService(ctrl)
			UserMock.EXPECT().Create(test.test).Return(test.want, test.err)
			res, err := UserMock.Create(test.test)
			tests.AssertEqual(t, res, test.want)
			tests.AssertEqual(t, err, test.err)
		})
	}
}
