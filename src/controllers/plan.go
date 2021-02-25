package controllers

import (
	"src/models/plan"
	"src/services/planservice"
    "net/http"
	"github.com/gin-gonic/gin"
)

type PlanIO struct {
    Name string `json:"name"`
    Insurance string `json:"insurance"`
    Mount float32 `json:"mount"`
    Mount_uf float32 `json":mount_uf"`
    Quota string `json:"quota"`
    Insured_capital float32 `json:"insured_capital"`
}

type PlanController interface {
	PostPlan(*gin.Context)
    GetPlan(*gin.Context)
}

type planController struct {
	ps planservice.PlanService
}

// NewPlanController instantiates Plan Controller
func NewPlanController(ps planservice.PlanService) PlanController {
	return &planController{ps: ps}
}

func (ctl *planController) PostPlan(c *gin.Context) {
	// Read user input
	var planInput PlanIO
	if err := c.ShouldBindJSON(&planInput); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	p := ctl.inputToPlan(planInput)

	// Create plan
	// If an Error Occurs while creating return the error
	if _, err := ctl.ps.CreatePlan(&p); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	planOutput := ctl.mapToPlanOutput(&p)
	HTTPRes(c, http.StatusOK, "Plan Published", planOutput)
}

func (ctl *planController) GetPlan(c *gin.Context){
    list,code, err := ctl.ps.GetPlan()
    if err != nil{
        panic(err)
    }
    response := map[string]interface{}{
        "plans": list,
    }
    c.JSON(code, response)
}

// Private Methods
func (ctl *planController) inputToPlan(input PlanIO) plan.Plan {
	return plan.Plan{
		Name:  input.Name,
        Insurance: input.Insurance,
        Mount: input.Mount,
        Mount_uf: input.Mount_uf,
        Quota: input.Quota,
        Insured_capital: input.Insured_capital,
	}
}
func (ctl *planController) mapToPlanOutput(p *plan.Plan) * PlanIO {
	return &PlanIO{
		Name:  p.Name,
        Insurance: p.Insurance,
        Mount: p.Mount,
        Mount_uf: p.Mount_uf,
        Quota: p.Quota,
        Insured_capital: p.Insured_capital,
	}
}
