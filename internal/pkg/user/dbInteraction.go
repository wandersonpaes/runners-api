package user

import (
	"database/sql"
	"log"
)

type userConnection struct {
	db *sql.DB
}

func newUserConnection(db *sql.DB) *userConnection {
	return &userConnection{db}
}

func (userTable userConnection) create(user User) (uint64, error) {
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
