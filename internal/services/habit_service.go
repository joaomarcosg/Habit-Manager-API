package services

import "github.com/joaomarcosg/Habit-Manager-API/internal/store"

type HabitService struct {
	Store store.HabitStore
}

func NewHabitService(store store.HabitStore) *HabitService {
	return &HabitService{
		Store: store,
	}
}
