
package controllers

import (
	"src/models/quotation"
	"src/services/quotationservice"
	"src/services/planservice"
    "net/http"
	"github.com/gin-gonic/gin"
)

type QuotationIO struct {
    Rut float64 `json:"rut"`
    Date_birth string `json:"date_birth"`
    Email string `json:"email"`
    Phone string `json:"phone"`
}

//QuotationController interface
type QuotationController interface {
	PostQuotation(*gin.Context)
}

type quotationController struct {
	qs quotationservice.QuotationService
    ps planservice.PlanService
}

// NewQuotationController instantiates Quotation Controller
func NewQuotationController(qs quotationservice.QuotationService, ps planservice.PlanService,
) QuotationController {
    return &quotationController{qs:qs, ps:ps}
}

func (ctl *quotationController) PostQuotation(c *gin.Context) {
	// Read user input
	var quotationInput QuotationIO
	if err := c.ShouldBindJSON(&quotationInput); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	q := ctl.inputToQuotation(quotationInput)

    // Create quotation
	// If an Error Occurs while creating return the error
	if _, err := ctl.qs.CreateQuotation(&q); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

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
func (ctl *quotationController) inputToQuotation(input QuotationIO) quotation.Quotation {
	return quotation.Quotation{
		Rut:  input.Rut,
        Date_birth: input.Date_birth,
        Email: input.Email,
        Phone: input.Phone,
	}
}
func (ctl *quotationController) mapToQuotationOutput(q *quotation.Quotation) * QuotationIO {
	return &QuotationIO{
		Rut:  q.Rut,
        Date_birth: q.Date_birth,
        Email: q.Email,
        Phone: q.Phone,
	}
}
