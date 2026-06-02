package service

import (
	"errors"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func AddCategory(role int, name string) error {
	if role != models.RoleAdmin {
		return errors.New("только администратор может добавлять категории")
	}
	if name == "" {
		return errors.New("название категории не может быть пустым")
	}
	return repository.AddCategory(name)
}

func GetCategories() ([]models.Category, error) {
	return repository.GetCategories()
}
