package main

import (
	"ShoppingCartAnalysis/internal/config"
	"ShoppingCartAnalysis/internal/handlers"
	"ShoppingCartAnalysis/internal/logger"
	"ShoppingCartAnalysis/internal/middleware"
	"context"

	"ShoppingCartAnalysis/internal/services"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	logger.InitLogger()
	cfg := config.Load()

	ctx := context.Background()
	
	// export LANGCHAINGO_TEST_POSTGRESQL=postgres://db_user:mysecretpassword@localhost:5438/test?sslmode=disable
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PG_User,
		cfg.PG_Password,
		cfg.PG_Host,
		cfg.PG_Port,
		cfg.PG_Dbname,
	)
	logger.LoggerInfo.Println("dsn: ", dsn)
	logger.LoggerInfo.Println(cfg.PG_User)

	geminiClient, err := googleai.New(ctx, googleai.WithAPIKey(cfg.GeminiAPIKey), googleai.WithDefaultModel("gemini-2.5-pro"))
	if err != nil {
		logger.LoggerError.Fatalln("Error Initializing LLM Client: ", err)
	}

	llmClient := services.NewGeminiService(dsn, geminiClient)

	r := gin.Default()
	r.Use(middleware.TracerMiddleWare())
	
	r.POST("/query", handlers.NewQueryHandler(llmClient).HandleQuery)

	logger.LoggerInfo.Printf("🚀 Server running on http://localhost:%s\n", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		logger.LoggerError.Fatal(err)
	}
}
