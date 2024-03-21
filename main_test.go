package main

import (
	"context"
	"dependency_injection/mocks"
	"dependency_injection/repository"
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_Main(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockUser := mocks.NewMockDatabase(mockCtl)
	gomock.InOrder(
		mockUser.EXPECT().GetUsers(context.Background()).Return([]repository.UserTable{
			{ID: 1, Name: "Sarthak"},
			{ID: 2, Name: "User1"},
		}, nil).Times(1),
	)

	user := User{
		name: "sarthak",
		db:   mockUser,
	}

	fmt.Println(user.db.GetUsers(context.Background()))
}
