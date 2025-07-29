package service

import (
    ""
)

func CreateUser(user *model.User) error {
    return repository.CreateUser(user)
}

func GetUsers() ([]model.User, error) {
    return repository.GetAllUsers()
}

func GetUser(id uint) (model.User, error) {
    return repository.GetUserByID(id)
}

func UpdateUser(user *model.User) error {
    return repository.UpdateUser(user)
}

func DeleteUser(id uint) error {
    return repository.DeleteUser(id)
}
