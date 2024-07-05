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

	//mapStringTodo := structs.Map(todo)

	//j, _ := json.Marshal(mapStringTodo)
	//fmt.Println(string(j))

	result := db.Model(&todo).Select("title", "done", "description", "deadline").Updates(&todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}
