package todos

import (
	"todolist/lib/gormLib"
	"todolist/models"
)

func GetAllTodos(userId uint, page int, perPage int) ([]models.Todo, error) {
	db := gormLib.CreateConnection()
	var todos []models.Todo

	db.Scopes(gormLib.Paginate(page, perPage)).Where(models.Todo{UserID: userId}).Find(&todos)

	return todos, nil
}

func CreateTodo(todo models.Todo) (*models.Todo, error) {

	db := gormLib.CreateConnection()

	result := db.Create(&todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func EditTodo(todo models.Todo) (*models.Todo, error) {
	db := gormLib.CreateConnection()

	var queryTodo models.Todo

	queryTodo.ID = todo.ID
	queryTodo.UserID = todo.UserID

	//check if not exists
	if result := db.Where(&queryTodo).First(&models.Todo{}); result.Error != nil {
		return nil, result.Error
	}


	result := db.Model(&models.Todo{}).Where(&queryTodo).Updates(&todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func DeleteTodo(todo models.Todo) (*models.Todo, error) {
	db := gormLib.CreateConnection()

	var queryTodo models.Todo

	queryTodo.ID = todo.ID
	queryTodo.UserID = todo.UserID

	//check if not exists
	if result := db.Where(&queryTodo).First(&models.Todo{}); result.Error != nil {
		return nil, result.Error
	}

	result := db.Model(&models.Todo{}).Where(&queryTodo).Delete(&todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}
