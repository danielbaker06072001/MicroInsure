package interfaces

import (
	application "PaymentService/Application"
	model "PaymentService/Model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service *application.PaymentService
}

func NewPaymentHandler(service *application.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}

func (h *PaymentHandler) CreatePayment(context *gin.Context) {
	var input struct {
		ID          int     `json:"id"`
		PaymentNumber string  `json:"Payment_number"`
		PaymentType   string  `json:"Payment_type"`
		PaymentAmount float64 `json:"Payment_amount"`
		PaymentDate   string  `json:"Payment_date"`
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// Parse the date manually
	parsedDate, err := time.Parse("2006-01-02", input.PaymentDate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid date format, expected YYYY-MM-DD"})
		return
	}

	// Convert to the model struct
	Payment := model.Payment{
		ID:          input.ID,
		PaymentNumber: input.PaymentNumber,
		PaymentType:   input.PaymentType,
		PaymentAmount: input.PaymentAmount,
		PaymentDate:   parsedDate, 
	}

	createdPayment, err := h.service.CreatePayment(&Payment)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "success", "Payment": createdPayment})
}


func (h *PaymentHandler) GetAllPayment(context *gin.Context) {
	Payment , err := h.service.GetAllPayment()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "success", "Payment" : Payment})
}

func (h *PaymentHandler) ProcessPayment(message string) {
	log.Printf("Processed Payment with ID: %s", message)
}