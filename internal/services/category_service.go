package services

import "github.com/joaomarcosg/Habit-Manager-API/internal/store"

type CategoryService struct {
	Store store.CategoryStore
}

func NewCategoryService(store store.CategoryStore) *CategoryService {
	return &CategoryService{
		Store: store,
	}
}
