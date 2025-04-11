package repository

import (
	"database/sql"
	"errors"
	"mps_notas_back/internal/infra/model"
)

type TaskRepositorySql struct {
	DB *sql.DB
}

var taskQueries = map[string]string{

	"create": " INSERT INTO tasks (title, description, start_date, end_date) VALUES (?, ?, ?, ?); ",

	"update": "UPDATE tasks SET title = ?, description = ?, start_date = ?, end_date = ?  WHERE id = ?;",

	"delete": " DELETE FROM tasks  WHERE id = ? ",

	"findAll": " SELECT id, title, description, start_date, end_date, created_at  FROM tasks t ",

	"findById": " SELECT id, title, description, start_date, end_date, created_at FROM tasks t WHERE t.id = ? ",
}

func (u TaskRepositorySql) FindAll() ([]model.TaskDAO, error) {
	//format the string to %% s %%

	rows, err := u.DB.Query(taskQueries["findAll"])
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []model.TaskDAO
	for rows.Next() {
		var task model.TaskDAO
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.StartDate, &task.EndDate, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (u TaskRepositorySql) FindByID(id int) (*model.TaskDAO, error) {
	rows, err := u.DB.Query(taskQueries["findById"], id)
	if err != nil {
		return &model.TaskDAO{}, err
	}

	defer rows.Close()
	var task model.TaskDAO
	if rows.Next() {
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.StartDate, &task.EndDate, &task.CreatedAt); err != nil {
			return &model.TaskDAO{}, err
		}
	}
	return &task, nil
}



func (u TaskRepositorySql) Create(input model.NewTaskInput) error {
	statement, err := u.DB.Prepare(taskQueries["create"])
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(input.Title, input.Description, input.StartDate, input.EndDate)
	return err
}

func (u TaskRepositorySql) Update(id int, input model.NewTaskInput) (model.TaskDAO, error) {
	// Verificar se o ID foi fornecido
	if id == 0 {
		return model.TaskDAO{}, errors.New("ID invalido")
	}

	// Executar a atualização
	stmt, err := u.DB.Prepare(taskQueries["update"])
	if err != nil {
		return model.TaskDAO{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(input.Title, input.Description, input.StartDate, input.EndDate, id)
	if err != nil {
		return model.TaskDAO{}, err
	}

	// Retornar o usuário atualizado
	updatedTask, err := u.FindByID(id)
	if err != nil {
		return model.TaskDAO{}, err
	}

	return model.TaskDAO{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		StartDate:   updatedTask.StartDate,
		EndDate:     updatedTask.EndDate,
		CreatedAt:   updatedTask.CreatedAt,
	}, nil
}

func (u TaskRepositorySql) Delete(id int) error {
	// Verificar se o ID é válido
	if id == 0 {
		return errors.New("ID invalido")
	}

	stmt, err := u.DB.Prepare(taskQueries["delete"])
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
