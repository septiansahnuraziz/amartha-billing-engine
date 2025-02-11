package http

import (
	"amartha-billing-engine/internal/entity"
	"amartha-billing-engine/utils"
	"amartha-billing-engine/utils/httpresponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Router) initLoansURLRoutes(app *gin.RouterGroup) {
	loans := app.Group("loans")
	{
		loans.POST("", r.createLoan)
		loans.POST("pay", r.payLoan)
		loans.GET("/:id", r.getLoan)
	}
}

// Endpoint Create Loan
//
//	@Summary	Endpoint for create loan
//	@Description
//	@Tags		Loans
//	@Accept		json
//	@Produce	json
//	@Param		Accept			header		string						false	"Example: application/json"
//	@Param		Content-Type	header		string						false	"Example: application/json"
//	@Param		request			body		entity.RequestCreateLoan	false	"Request Body"
//	@Success	200				{object}	entity.SwaggerResponseOKDTO{}
//	@Router		/v1/loans [post]
func (r *Router) createLoan(context *gin.Context) {
	var request entity.RequestCreateLoan
	if err := context.Bind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := r.LoanService.CreateLoan(context, request); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "succes create loan"})
	return
}

// Endpoint Create Pay Loan
//
//	@Summary	Endpoint for create pay loan
//	@Description
//	@Tags		Loans
//	@Accept		json
//	@Produce	json
//	@Param		Accept			header		string						false	"Example: application/json"
//	@Param		Content-Type	header		string						false	"Example: application/json"
//	@Param		request			body		entity.RequestPayLoan	false	"Request Body"
//	@Success	200				{object}	entity.SwaggerResponseOKDTO{}
//	@Router		/v1/loans/pay [post]
func (r *Router) payLoan(context *gin.Context) {
	var request entity.RequestPayLoan
	if err := context.Bind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := r.LoanService.PayLoan(context, request); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "succes create loan"})
	return
}

// Endpoint Get Loan Detail
//
//	@Summary	Endpoint for get loan detail
//	@Description
//	@Tags		Loans
//	@Accept		json
//	@Produce	json
//	@Param		Accept			header		string						false	"Example: application/json"
//	@Param		Content-Type	header		string						false	"Example: application/json"
//	@Param		id				path		string						true	"Loan Id"
//	@Success	200				{object}	entity.ResponseGetLoanDetail{}
//	@Router		/v1/loans/{id} [get]
func (r *Router) getLoan(context *gin.Context) {
	loanID := context.Param("id")

	existLoan, err := r.LoanService.GetLoanDetail(context, utils.ExpectedUint(loanID))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	httpresponse.NewHttpResponse().WithData(existLoan).WithMessage("Success Get Loan Detail").ToWrapperResponseDTO(context, http.StatusOK)
}
