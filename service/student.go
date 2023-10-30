package service

import (
	"go-echo-gorm-practice/model"
	"go-echo-gorm-practice/storage"
)

func Students() ([]model.Student, error) {
	db := storage.GetDB()
	students := []model.Student{}

	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}