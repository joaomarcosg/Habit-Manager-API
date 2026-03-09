package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
)

type CategoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (cs *CategoryService) CreateCategory(ctx context.Context, name string) (uuid.UUID, error) {

	category := domain.Category{
		Name: name,
	}

	id, err := cs.repo.CreateCategory(ctx, category)

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (cs *CategoryService) GetCategoryByName(ctx context.Context, name string) (domain.Category, error) {

	category, err := cs.repo.GetCategoryByName(ctx, name)

	if err != nil {
		return domain.Category{}, err
	}

	return category, nil

}

func (cs *CategoryService) GetCategoryEntries(ctx context.Context, name string) (domain.Category, error) {

	categoryEntries, err := cs.repo.GetCategoryByName(ctx, name)
	if err != nil {
		return domain.Category{Entries: 0}, err
	}

	return domain.Category{
		Entries: categoryEntries.Entries,
	}, nil
}

func (cs *CategoryService) DeleteCategory(ctx context.Context, name string) (bool, error) {

	category, err := cs.repo.GetCategoryByName(ctx, name)

	if err != nil {
		return false, err
	}

	ok, err := cs.repo.DeleteCategory(ctx, category.Name)

	if err != nil {
		return false, err
	}

	return ok, nil

}
