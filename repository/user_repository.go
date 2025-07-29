package repository

import (
    "go-fiber-crud/config"
    "go-fiber-crud/model"
)

func CreateUser(user *model.User) error {
    return config.DB.Create(user).Error
}

func GetAllUsers() ([]model.User, error) {
    var users []model.User
    err := config.DB.Find(&users).Error
    return users, err
}

func GetUserByID(id uint) (model.User, error) {
    var user model.User
    err := config.DB.First(&user, id).Error
    return user, err
}

func UpdateUser(user *model.User) error {
    return config.DB.Save(user).Error
}

func DeleteUser(id uint) error {
    return config.DB.Delete(&model.User{}, id).Error
}
