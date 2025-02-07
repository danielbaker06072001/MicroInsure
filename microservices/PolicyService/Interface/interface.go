package interfaces

import (
	application "PolicyService/Application"
	model "PolicyService/Model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PolicyHandler struct {
	service *application.PolicyService
}

func NewPolicyHandler(service *application.PolicyService) *PolicyHandler {
	return &PolicyHandler{
		service: service,
	}
}

func (h *PolicyHandler) CreatePolicy(context *gin.Context) {
	var input struct {
		ID          int     `json:"id"`
		PolicyNumber string  `json:"Policy_number"`
		PolicyType   string  `json:"Policy_type"`
		PolicyAmount float64 `json:"Policy_amount"`
		PolicyDate   string  `json:"Policy_date"`
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// Parse the date manually
	parsedDate, err := time.Parse("2006-01-02", input.PolicyDate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid date format, expected YYYY-MM-DD"})
		return
	}

	// Convert to the model struct
	Policy := model.Policy{
		ID:          input.ID,
		PolicyNumber: input.PolicyNumber,
		PolicyType:   input.PolicyType,
		PolicyAmount: input.PolicyAmount,
		PolicyDate:   parsedDate, 
	}

	createdPolicy, err := h.service.CreatePolicy(&Policy)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "success", "Policy": createdPolicy})
}


func (h *PolicyHandler) GetAllPolicy(context *gin.Context) {
	Policy , err := h.service.GetAllPolicy()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "success", "Policy" : Policy})
}