package service

import (
	"errors"
	"go-test/entity"
	"go-test/repository"
)

type CategoryService struct {
	Rc repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	categoryck := service.Rc.FindById(id)
	if categoryck == nil {
		return categoryck, errors.New("category not found")
	} else {
		return categoryck, nil
	}

}
