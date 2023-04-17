package service

import "log"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Save() {
	log.Println("用户保存成功")
}
