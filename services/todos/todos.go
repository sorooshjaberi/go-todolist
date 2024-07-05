package todos

import (
	"todolist/lib/gormLib"
	"todolist/models"
)

func GetAllTodos(userId uint, page int, perPage int) ([]models.Todo, error) {
	db := gormLib.CreateConnection()
	var todos []models.Todo

	db.Scopes(gormLib.Paginate(page, perPage)).Where(models.Todo{UserID: 6}).Find(&todos)

	return todos, nil
}

func CreateTodo(userId uint, todo models.Todo) (*models.Todo, error) {
	todo.UserID = userId

	db := gormLib.CreateConnection()

	result := db.Create(&todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}
