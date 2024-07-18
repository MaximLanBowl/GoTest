package repository

import (
	"database/sql"
	"my-microservice/internal/model"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository() *PostgresRepository {
	db, err := sql.Open("postgres", "user=myuser password=mypassword dbname=mydb sslmode=disable")
	if err != nil {
		panic(err)
	}

	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) SaveMessage(msg *model.Message) error {
	_, err := r.db.Exec("INSERT INTO messages (text, processed) VALUES ($1, $2)", msg.Text, msg.Processed)
	return err
}

func (r *PostgresRepository) GetMessages() ([]*model.Message, error) {
	rows, err := r.db.Query("SELECT id, text, processed FROM messages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*model.Message
	for rows.Next() {
		var msg model.Message
		if err := rows.Scan(&msg.ID, &msg.Text, &msg.Processed); err != nil {
			return nil, err
		}
		messages = append(messages, &msg)
	}

	return messages, nil
}