package repository

import (
	"database/sql"
	"errors"
	"mps_notas_back/internal/buisness/security"
	"mps_notas_back/internal/infra/model"
)

type UserRepositorySql struct {
	DB *sql.DB
}

var userQueries = map[string]string{

	"create": " INSERT INTO users (name, email, password_hash) VALUES (?, ?, ?); ",

	"update": "UPDATE users SET name = ?, email = ?, password_hash = ?  WHERE id = ?;",

	"delete": " DELETE FROM users  WHERE id = ? ",

	"findAll": " SELECT id, name, email, created_at FROM users u ", //format the string to %% s %%

	"findById": " SELECT id, name ,email, created_at FROM users u WHERE u.id = ? ",

	"findByEmail": " SELECT id, name ,email, created_at FROM users u WHERE u.email = ? ",
}

func (u UserRepositorySql) FindAll() ([]model.UserDAO, error) {
	//format the string to %% s %%

	rows, err := u.DB.Query(userQueries["findAll"])
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.UserDAO
	for rows.Next() {
		var user model.UserDAO
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u UserRepositorySql) FindByID(id int) (*model.UserDAO, error) {
	rows, err := u.DB.Query(userQueries["findById"], id)
	if err != nil {
		return &model.UserDAO{}, err
	}

	defer rows.Close()
	var user model.UserDAO
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return &model.UserDAO{}, err
		}
	}
	return &user, nil
}

func (u UserRepositorySql) FindByEmail(email string) (*model.UserDAO, error) {

	rows, err := u.DB.Query(userQueries["findByEmail"], email)
	if err != nil {
		return &model.UserDAO{}, err
	}

	defer rows.Close()
	var user model.UserDAO
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return &model.UserDAO{}, err
		}
	}
	return &user, nil

}

func (u UserRepositorySql) Create(input model.NewUserInput) (model.UserDAO, error) {
	statement, err := u.DB.Prepare(userQueries["create"])
	if err != nil {
		return model.UserDAO{}, err
	}
	defer statement.Close()

	hashedBytes, err := security.Hash(input.Password)
	if err != nil {
		return model.UserDAO{}, err
	}
	input.Password = string(hashedBytes)

	_, err = statement.Exec(input.Name, input.Email, input.Password)
	if err != nil {
		return model.UserDAO{}, err
	}
	row, err := u.DB.Query(userQueries["findByEmail"], input.Email)
	if err != nil {
		return model.UserDAO{}, err
	}
	defer row.Close()
	var userToReturn model.UserDAO
	if row.Next() {
		if err := row.Scan(
			&userToReturn.ID,
			&userToReturn.Name,
			&userToReturn.Email,
			&userToReturn.CreatedAt,
		); err != nil {
			return model.UserDAO{}, err
		}
	}
	return userToReturn, nil
}

func (u UserRepositorySql) Update(id int, input model.NewUserInput) (model.UserDAO, error) {
	// Verificar se o ID foi fornecido
	if id == 0 {
		return model.UserDAO{}, errors.New("ID invalido")
	}

	hashedBytes, err := security.Hash(input.Password)
	if err != nil {
		return model.UserDAO{}, err
	}
	input.Password = string(hashedBytes)

	// Hash da senha se fornecida
	if input.Password != "" {
		hashedBytes, err := security.Hash(input.Password)
		if err != nil {
			return model.UserDAO{}, err
		}
		input.Password = string(hashedBytes)
	}

	// Executar a atualização
	stmt, err := u.DB.Prepare(userQueries["update"])
	if err != nil {
		return model.UserDAO{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(input.Name, input.Email, input.Password, id)
	if err != nil {
		return model.UserDAO{}, err
	}

	// Retornar o usuário atualizado
	updatedUser, err := u.FindByID(id)
	if err != nil {
		return model.UserDAO{}, err
	}

	return model.UserDAO{
		ID:        updatedUser.ID,
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		Password_Hash:     updatedUser.Password_Hash,
		CreatedAt: updatedUser.CreatedAt,
	}, nil
}

func (u UserRepositorySql) Delete(id int) error {
	// Verificar se o ID é válido
	if id == 0 {
		return errors.New("ID invalido")
	}

	stmt, err := u.DB.Prepare(userQueries["delete"])
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
