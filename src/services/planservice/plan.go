package planservice

import (
    "src/models/plan"
	"src/repositories/planrepo"
)

type PlanService interface {
	CreatePlan(plan *plan.Plan) (*plan.Plan, error)
    GetPlan() ([]plan.Plan, int, error)
}

type planService struct {
	Repo planrepo.Repo
}

// NewPlanService will instantiate User Service
func NewPlanService(
	repo planrepo.Repo,
) PlanService {

	return &planService{
		Repo: repo,
	}
}

func (ps *planService) CreatePlan(plan *plan.Plan) (*plan.Plan, error) {
	return ps.Repo.CreatePlan(plan)
}

func (ps *planService) GetPlan() ([]plan.Plan, int, error){
    return ps.Repo.GetPlan()
}
