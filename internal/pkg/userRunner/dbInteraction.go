package userRunner

import (
	"database/sql"
	"fmt"
	"log"
)

type UserConnection struct {
	db *sql.DB
}

func NewUserConnection(db *sql.DB) *UserConnection {
	return &UserConnection{db}
}

func (userTable UserConnection) create(user User) (uint64, error) {
	statement, err := userTable.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return uint64(id), nil
}

func (userTable UserConnection) search(nameOrNick string) ([]User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	line, err := userTable.db.Query(
		"select id, name, nick, email, createOn from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}
	defer line.Close()

	var users []User

	for line.Next() {
		var user User

		if err = line.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateOn,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (userTable UserConnection) searchByID(ID uint64) (User, error) {
	lines, err := userTable.db.Query(
		"select id, name, nick, email, createOn from users where id = ?",
		ID,
	)
	if err != nil {
		return User{}, err
	}
	defer lines.Close()

	var user User
	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateOn,
		); err != nil {
			return User{}, err
		}
	}

	return user, nil
}

func (userTable UserConnection) update(ID uint64, user User) error {
	statement, err := userTable.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (userTable UserConnection) delete(ID uint64) error {
	statement, err := userTable.db.Prepare(
		"delete from users where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (userTable UserConnection) SearchByEmail(email string) (User, error) {
	line, err := userTable.db.Query(
		"select id, password from users where email = ?",
		email,
	)
	if err != nil {
		return User{}, err
	}
	defer line.Close()

	var user User
	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return User{}, err
		}
	}

	return user, nil
}
