package postgres

import (
	"log"
	"os"
	"skillfactory/go_news/pkg/storage"
	"testing"
)

var s *Store

func TestMain(m *testing.M) {
	connstr := "postgres://postgres:123456@localhost:5432/posts"
	var err error
	s, err = New(connstr)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestStorage_Posts(t *testing.T) {
	data, err := s.Posts()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestStorage_AddPost(t *testing.T) {
	task := storage.Post{
		Title: "unit test task title",
		Content: "unit test task content",
		AuthorID: 0,
		CreatedAt: 11231231323,
	}
	err := s.AddPost(task)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Создан пост:")
}

func TestStorage_DeletePost(t *testing.T) {
	err := s.DeletePost(storage.Post{ID: 2})
	if err != nil {
		t.Fatal(err)
	}
}

func TestStorage_UpdatePost(t *testing.T) {
	task := storage.Post{
		ID: 2,
		Title: "unit test task title",
		Content: "unit test task content",
		AuthorID: 0,
		CreatedAt: 212312312312,
	}
	err := s.UpdatePost(task)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Изменен пост")
}