package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"skillfactory/go_news/pkg/storage"
)

const (
	databaseName = "posts"
	collectionName = "posts"
)

// Хранилище данных.
type Store struct {
	db *mongo.Client
}

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Store, error) {
	mongoOpts := options.Client().ApplyURI(constr)
	client, err := mongo.Connect(context.Background(), mongoOpts)
	fmt.Println(client)
	if err != nil {
		log.Fatal(err)
	}
	// проверка связи с БД
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	s := Store{
		db: client,
	}
	return &s,nil
}

// Tasks возвращает список постов
func (s *Store) Posts() ([]storage.Post, error) {
	collection := s.db.Database(databaseName).Collection(collectionName)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var data []storage.Post
	for cur.Next(context.Background()) {
		var l storage.Post
		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}
		data = append(data, l)
	}
	return data, cur.Err()
}

// AddPost создаёт новый пост
func (s *Store) AddPost(p storage.Post) (error) {
	collection := s.db.Database(databaseName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), p)
	return err
}

// UpdateTask обновляет пост
func (s *Store) UpdatePost(p storage.Post) (error) {
	collection := s.db.Database(databaseName).Collection(collectionName)
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"id", p.ID}}
	update := bson.D{{"$set", p}}

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// DeletePost удаляет по её id
func (s *Store) DeletePost(p storage.Post) (error) {
	collection := s.db.Database(databaseName).Collection(collectionName)
	_, err := collection.DeleteOne(context.TODO(), bson.D{{"id", p.ID}}, options.Delete().SetCollation(&options.Collation{}))
	if err != nil {
		log.Fatal(err)
	}
	return err
}


