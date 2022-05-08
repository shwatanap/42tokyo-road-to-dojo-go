package repository

import (
	"context"
	"database/sql"
	"log"

	"42tokyo-road-to-dojo-go/pkg/domain/entity"
	"42tokyo-road-to-dojo-go/pkg/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Create(ctx context.Context, name string) (*entity.User, error) {
	const (
		insert  = `INSERT INTO users (name, token, high_score, coin) VALUE (?, UUID(), 0, 0)`
		confirm = `SELECT token FROM users WHERE id = ?`
	)

	stmt, err := ur.db.PrepareContext(ctx, insert)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, name)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	ue := &entity.User{}

	stmt, err = ur.db.PrepareContext(ctx, confirm)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRowContext(ctx, id).Scan(&ue.Token)
	if err != nil {
		return nil, err
	}

	return ue, nil
}

func (ur *userRepository) Get(ctx context.Context, token string) (*entity.User, error) {
	const read = `SELECT name FROM users WHERE token = ?`

	stmt, err := ur.db.PrepareContext(ctx, read)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, token)
	if err = row.Err(); err != nil {
		log.Println("Exec", err.Error())
		return nil, err
	}

	ue := &entity.User{}
	err = row.Scan(&ue.Name)
	if err != nil {
		log.Println("Scan", err.Error())
		return nil, err
	}

	return ue, nil
}
