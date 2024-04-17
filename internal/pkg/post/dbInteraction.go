package post

import (
	"database/sql"
)

type postConnection struct {
	db *sql.DB
}

func newPostConnection(db *sql.DB) *postConnection {
	return &postConnection{db}
}

func (postTable postConnection) create(post Posts) (uint64, error) {
	statement, err := postTable.db.Prepare(
		"insert into posts (title, postText, author_id) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.PostText, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

func (postTable postConnection) searchByID(postID uint64) (Posts, error) {
	line, err := postTable.db.Query(`
		select p.*, u.nick from
		posts p inner join users u
		on u.id = p.author_id where p.id = ?`,
		postID,
	)
	if err != nil {
		return Posts{}, err
	}
	defer line.Close()

	var post Posts
	if line.Next() {
		if err = line.Scan(
			&post.ID,
			&post.Title,
			&post.PostText,
			&post.AuthorID,
			&post.Likes,
			&post.CreateOn,
			&post.AuthorNick,
		); err != nil {
			return Posts{}, err
		}
	}

	return post, nil
}