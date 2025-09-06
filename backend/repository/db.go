package repository

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"todolist/model"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() Repository {
	db, err := gorm.Open(sqlite.Open("todolist"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	// Migrate the schema
	if err := db.AutoMigrate(&model.TodoItem{}); err != nil {
		panic("failed to migrate database")
	}
	
	return Repository{db: db}
}

func (repo *Repository) GetAll() ([]*model.TodoItem, error) {
	var items []*model.TodoItem
	err := repo.db.Find(&items).Error
	return items, err
}

func (repo *Repository) Add(item *model.TodoItem) error {
	return repo.db.Create(item).Error
}
func (repo *Repository) Update(item *model.TodoItem) error {
	fmt.Println(item)
	return repo.db.Save(item).Error
}

func (repo *Repository) Delete(id string) error {
	return repo.db.Where("id = ?", id).Delete(&model.TodoItem{}).Error
}
