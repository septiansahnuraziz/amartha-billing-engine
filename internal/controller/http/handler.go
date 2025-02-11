package http

import (
	"amartha-billing-engine/internal/entity"
	"github.com/gin-gonic/gin"
)

type Router struct {
	LoanService entity.ILoanService
}

func NewRoute(app *gin.RouterGroup, loanService entity.ILoanService) {
	router := &Router{
		LoanService: loanService,
	}

	router.handlers(app)
}

func (r *Router) handlers(app *gin.RouterGroup) {
	//app.GET("/ping", ping)

	apiGroupV1 := app.Group("v1")
	{
		r.initLoansURLRoutes(apiGroupV1)
	}
}

//func ping(c *gin.Context) {
//	response := httpresponse.NewHttpResponse()
//	httpresponse.NoContent(c, response)
//}
