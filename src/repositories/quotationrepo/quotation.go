package quotationrepo

import (
	"context"
	"src/models/quotation"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repo interface {
	CreateQuotation(quotation *quotation.Quotation) (*quotation.Quotation, error)
}

type quotationRepo struct {
	db *mongo.Client
}

// NewQuotationRepo will instantiate Plan Repository
func NewQuotationRepo(db *mongo.Client) Repo {
	return &quotationRepo{
		db: db,
	}
}

func (q *quotationRepo) CreateQuotation(quotation *quotation.Quotation) (*quotation.Quotation, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	collection := q.db.Database("quotations-db").Collection("quotations")
	_, err := collection.InsertOne(ctx, *quotation)

	if err != nil {
		panic(err)
	}
	return quotation, nil
}
