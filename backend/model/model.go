package model

type TodoItem struct {
	ID        string `gorm:"type:text;primaryKey" json:"id"`
	Content   string `gorm:"type:text" json:"content"`
	Completed bool   `gorm:"type:integer" json:"isCompleted"`
}
