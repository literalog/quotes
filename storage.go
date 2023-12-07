package main

import (
	"context"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage interface {
	CreateQuote(*Quote) error
	GetQuoteById(id string) (*Quote, error)
	GetQuotes() ([]*Quote, error)
	UpdateQuote(q *Quote) error
	DeleteQuote(id string) error
}

type MongoStorage struct {
	client *mongo.Client
}

func NewMongoStorage() (*MongoStorage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		mongoUri = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}

	return &MongoStorage{
		client: client,
	}, nil
}

func (s *MongoStorage) Init() error {
	return nil
}

func (s *MongoStorage) CreateQuote(q *Quote) error {
	col := s.client.Database("quotesdb").Collection("quotes")
	_, err := col.InsertOne(context.Background(), q)
	if err != nil {
		return fmt.Errorf("failed to insert quote: %w", err)
	}

	return nil
}

func (s *MongoStorage) GetQuoteById(id string) (*Quote, error) {
	var q Quote

	col := s.client.Database("quotesdb").Collection("quotes")
	if err := col.FindOne(context.Background(), bson.M{"_id": id}).Decode(&q); err != nil {
		return nil, fmt.Errorf("failed to find quote: %w", err)
	}

	return &q, nil
}

func (s *MongoStorage) GetQuotes() ([]*Quote, error) {
	var qq []*Quote

	col := s.client.Database("quotesdb").Collection("quotes")
	cur, err := col.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find quotes: %w", err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var q Quote
		if err := cur.Decode(&q); err != nil {
			return nil, fmt.Errorf("failed to decode quote: %w", err)
		}

		qq = append(qq, &q)
	}

	return qq, nil
}

func (s *MongoStorage) UpdateQuote(q *Quote) error {
	col := s.client.Database("quotesdb").Collection("quotes")
	_, err := col.UpdateOne(context.Background(), bson.M{"_id": q.Id}, bson.M{"$set": q})
	if err != nil {
		return fmt.Errorf("failed to update quote: %w", err)
	}

	return err
}

func (s *MongoStorage) DeleteQuote(id string) error {
	col := s.client.Database("quotesdb").Collection("quotes")
	_, err := col.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("failed to delete quote: %w", err)
	}

	return err
}
