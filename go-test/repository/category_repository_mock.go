package repository

import (
	"go-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (c *CategoryRepositoryMock) FindById(id string) *entity.Category {
	argument := c.Mock.Called(id)
	if argument.Get(0) == nil {
		return nil
	} else {
		category := argument.Get(0).(entity.Category)
		return &category
	}
}
