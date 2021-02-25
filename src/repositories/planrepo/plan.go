package planrepo

import (
	"context"
	"src/models/plan"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)

type Repo interface {
	CreatePlan(plan *plan.Plan) (*plan.Plan, error)
    GetPlan() ([]plan.Plan, int, error)
}

type planRepo struct {
	db *mongo.Client
}

// NewPlanRepo will instantiate Plan Repository
func NewPlanRepo(db *mongo.Client) Repo {
	return &planRepo{
		db: db,
	}
}

func (p *planRepo) CreatePlan(plan *plan.Plan) (*plan.Plan, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	collection := p.db.Database("plans-db").Collection("plans")
	_, err := collection.InsertOne(ctx, *plan)

	if err != nil {
		panic(err)
	}
	return plan, nil
}

func (p *planRepo) GetPlan() ([]plan.Plan, int, error) {
    planList := []plan.Plan{}
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    collection := p.db.Database("plans-db").Collection("plans")
    cursor, err := collection.Find(ctx, bson.M{})//bson.D{{}})

    if err != nil {
        return []plan.Plan{}, 400, err
    }

    for cursor.Next(ctx) {
        var plans plan.Plan
        err = cursor.Decode(&plans)
        if err != nil {
            panic(err)
        }
        planList = append(planList, plans)
    }
    return planList, 400, nil
}
