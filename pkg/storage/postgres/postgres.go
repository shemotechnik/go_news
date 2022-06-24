package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"skillfactory/go_news/pkg/storage"
)

// Хранилище данных.
type Store struct {
	db *pgxpool.Pool
}

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Store, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Store{
		db: db,
	}
	return &s, nil
}

// Tasks возвращает список постов
func (s *Store) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			author_id,
			title,
			content,
			created_at
		FROM posts
		ORDER BY id;
	`,
	)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t storage.Post
		err = rows.Scan(
			&t.ID,
			&t.AuthorID,
			&t.Title,
			&t.Content,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		posts = append(posts, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return posts, rows.Err()
}

// AddPost создаёт новый пост
func (s *Store) AddPost(p storage.Post) (error) {
	_,err := s.db.Exec(context.Background(), `
		INSERT INTO posts (title, content, author_id, created_at)
		VALUES ($1, $2, $3, $4);
		`,
		p.Title,
		p.Content,
		p.AuthorID,
		p.CreatedAt,
	)
	return err
}

// UpdateTask обновляет пост
func (s *Store) UpdatePost(p storage.Post) (error) {
	_,err := s.db.Exec(context.Background(), `
		UPDATE posts SET title=$1, content=$2, author_id=$3, created_at=$4
		WHERE id = $5;
		`,
		p.Title,
		p.Content,
		p.AuthorID,
		p.CreatedAt,
		p.ID,
	)
	return err
}

// DeletePost удаляет по её id
func (s *Store) DeletePost(p storage.Post) (error) {
	_,err := s.db.Exec(context.Background(), `
		DELETE FROM posts WHERE id=$1;
		`,
		p.ID,
	)
	return err
}
