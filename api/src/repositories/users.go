package repositories

import (
	"database/sql"

	"github.com/johnnyseubert/devbook/src/models"
)

type usersRepository struct {
	db *sql.DB
}

func UsersRepository(db *sql.DB) *usersRepository {
	return &usersRepository{db: db}
}

func (repository usersRepository) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (repository usersRepository) GetAll(userOrNickname string) ([]models.User, error) {
	userOrNickname = "%" + userOrNickname + "%"

	rows, err := repository.db.Query("SELECT id, name, nick, email, created_at FROM users WHERE name LIKE ? OR nick LIKE ?", userOrNickname, userOrNickname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersList []models.User
	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}

	return usersList, nil
}

func (repository usersRepository) GetById(id string) (models.User, error) {

	rows, err := repository.db.Query("SELECT id, name, nick, email, created_at FROM users WHERE id = ?", id)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)
		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository usersRepository) Update(id string, user models.User) error {
	statement, err := repository.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nick, user.Email, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository usersRepository) Delete(id string) error {
	statement, err := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (repository usersRepository) GetByEmail(email string) (models.User, error) {
	rows, err := repository.db.Query("SELECT id, name, nick, email, created_at, password FROM users WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt, &user.Password)
		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
