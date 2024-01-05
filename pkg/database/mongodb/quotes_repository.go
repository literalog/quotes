package mongodb

import (
	"context"
	"fmt"

	"github.com/literalog/quotes/pkg/quote"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuoteRepository struct {
	collection *mongo.Collection
}

func NewQuoteRepository(collection *mongo.Collection) quote.Repository {
	return &QuoteRepository{
		collection: collection,
	}
}

func (r *QuoteRepository) Create(ctx context.Context, q *quote.Quote) error {
	_, err := r.collection.InsertOne(ctx, q)
	if err != nil {
		return fmt.Errorf("error creating quote: %w", err)
	}
	return nil
}

func (r *QuoteRepository) Update(ctx context.Context, q *quote.Quote) error {
	filter := bson.M{"_id": q.Id}
	update := bson.M{"$set": q}
	if _, err := r.collection.UpdateOne(ctx, filter, update); err != nil {
		return fmt.Errorf("error updating quote: %w", err)
	}
	return nil
}

func (r *QuoteRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return fmt.Errorf("error deleting quote: %w", err)
	}
	return nil
}

func (r *QuoteRepository) GetById(ctx context.Context, id string) (*quote.Quote, error) {
	filter := bson.M{"_id": id}
	q := new(quote.Quote)
	if err := r.collection.FindOne(ctx, filter).Decode(q); err != nil {
		return nil, fmt.Errorf("error getting quote: %w", err)
	}
	return q, nil
}

func (r *QuoteRepository) GetAll(ctx context.Context) ([]quote.Quote, error) {
	qq := make([]quote.Quote, 0)
	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting quote: %w", err)
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &qq); err != nil {
		return nil, fmt.Errorf("error getting quote: %w", err)
	}

	return qq, nil
}
