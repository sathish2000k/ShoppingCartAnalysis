package handlers

import (
	"ShoppingCartAnalysis/internal/logger"
	"ShoppingCartAnalysis/internal/middleware"
	"ShoppingCartAnalysis/internal/models"
	"ShoppingCartAnalysis/internal/services"
	"ShoppingCartAnalysis/internal/utils"
	"context"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryHandler struct {
	llmClient *services.GeminiService
}

func NewQueryHandler(gemini *services.GeminiService) *QueryHandler {
	return &QueryHandler{llmClient: gemini}
}

func (queryHandler *QueryHandler) HandleQuery(c *gin.Context) {
	var request models.QueryRequest

	traceID, _ := c.Get(middleware.TraceIDKey)
	ctx := utils.WithTraceID(context.Background(), traceID.(string))

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	logger.LoggerInfo.Printf("[%s] Received query: %s", traceID, request)

	answer, err := queryHandler.llmClient.RunGeminiQuery(ctx, request.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, models.QueryResponse{Answer: answer})
}