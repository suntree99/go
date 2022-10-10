package service

import "03p_belajar_golang_unit_test/repository"
import "03p_belajar_golang_unit_test/entity"
import "errors"

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category not found")
	} else {
		return category, nil
	}
}