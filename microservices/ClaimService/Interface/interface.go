package interfaces

import (
	Config "ClaimService/AppConfig"
	application "ClaimService/Application"
	model "ClaimService/Model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ClaimHandler struct {
	service *application.ClaimService
}

func NewClaimHandler(service *application.ClaimService) *ClaimHandler {
	return &ClaimHandler{
		service: service,
	}
}

func (h *ClaimHandler) CreateClaim(context *gin.Context) {
	var input struct {
		ID          int     `json:"id"`
		ClaimNumber string  `json:"claim_number"`
		ClaimType   string  `json:"claim_type"`
		ClaimAmount float64 `json:"claim_amount"`
		ClaimDate   string  `json:"claim_date"`
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// Parse the date manually
	parsedDate, err := time.Parse("2006-01-02", input.ClaimDate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid date format, expected YYYY-MM-DD"})
		return
	}

	// Convert to the model struct
	claim := model.Claim{
		ID:          input.ID,
		ClaimNumber: input.ClaimNumber,
		ClaimType:   input.ClaimType,
		ClaimAmount: input.ClaimAmount,
		ClaimDate:   parsedDate, 
	}

	createdClaim, err := h.service.CreateClaim(&claim)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "success", "claim": createdClaim})
}


func (h *ClaimHandler) GetAllClaim(context *gin.Context) {
	fmt.Println(Config.DiscoverServiceWithConsul("payment-service-8081"))
	claim , err := h.service.GetAllClaim()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "success", "claim" : claim})
}