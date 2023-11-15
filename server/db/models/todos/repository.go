package todos

import (
	"gorm.io/gorm"
)

// TodoRepository is the interface that represents operations to interact with the todos' datastore.
type TodoRepository interface {
	Create(todo *Todo) (*Todo, error)
	FindByID(id uint) (*Todo, error)
	FindAll() ([]*Todo, error)
	FindAllByUserID(userID uint) ([]*Todo, error)
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) Create(todo *Todo) (*Todo, error) {
	result := r.db.Create(todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return todo, nil
}

func (r *todoRepository) FindAll() ([]*Todo, error) {
	var todos []*Todo
	result := r.db.Find(&todos)
	return todos, result.Error
}

func (r *todoRepository) FindByID(id uint) (*Todo, error) {
	var todo Todo
	result := r.db.First(&todo, id)
	return &todo, result.Error
}

func (r *todoRepository) FindAllByUserID(userID uint) ([]*Todo, error) {
	var todos []*Todo
	result := r.db.Where("user_id = ?", userID).Find(&todos)
	return todos, result.Error
}
