package quote

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, q *Quote) error
	Update(ctx context.Context, q *Quote) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]Quote, error)
	GetById(ctx context.Context, id string) (*Quote, error)
}

type mongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(col *mongo.Collection) Repository {
	return &mongoRepository{
		collection: col,
	}
}

func (r *mongoRepository) Create(ctx context.Context, q *Quote) error {
	_, err := r.collection.InsertOne(ctx, q)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoRepository) Update(ctx context.Context, q *Quote) error {
	filter := bson.M{"_id": q.Id}
	update := bson.M{"$set": q}
	if _, err := r.collection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	return nil
}

func (r *mongoRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}

func (r *mongoRepository) GetAll(ctx context.Context) ([]Quote, error) {
	var qq []Quote
	cur, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &qq); err != nil {
		return nil, err
	}

	return qq, nil
}

func (r *mongoRepository) GetById(ctx context.Context, id string) (*Quote, error) {
	filter := bson.M{"_id": id}
	var q Quote
	if err := r.collection.FindOne(ctx, filter).Decode(&q); err != nil {
		return nil, err
	}
	return &q, nil
}
