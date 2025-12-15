package services

import (
	"ShoppingCartAnalysis/internal/logger"
	"ShoppingCartAnalysis/internal/utils"
	"context"
	"log"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/tools/sqldatabase"
	"github.com/tmc/langchaingo/tools/sqldatabase/postgresql"
)

type GeminiService struct {
	db *sqldatabase.SQLDatabase
	llm llms.Model
}

func NewGeminiService(dbDSN string, llm llms.Model) *GeminiService {
	engine, err := postgresql.NewPostgreSQL(dbDSN)
	if err != nil {
		logger.LoggerError.Fatalln("Failed to load engine: ", err)
	}

	db, err := sqldatabase.NewSQLDatabase(engine, nil)
	if err != nil {
		logger.LoggerError.Fatalln("Failed to connect sql database: ", err)
	}

	return &GeminiService{
		db: db,
		llm: llm,
	}
}

func (geminiService *GeminiService) RunGeminiQuery(ctx context.Context, query string) (string, error){
	traceID := utils.GetTraceID(ctx)

	tables := geminiService.db.TableNames()
	log.Printf("[%s] Printing the tables: %s", traceID, tables)

	sqlChain := chains.NewSQLDatabaseChain(geminiService.llm, 5, geminiService.db)

	input := map[string]interface{} {
		"query": query,
	}

	result, err := sqlChain.Call(ctx, input)
	if err != nil {
		logger.LoggerInfo.Println(traceID, "Error calling chain: ",  err)
	}

	logger.LoggerInfo.Printf("[%s] Response from LLM: %s", traceID, result["result"])

	return result["result"].(string), nil
}
