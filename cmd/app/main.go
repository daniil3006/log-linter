package main

import (
	"context"
	"log-linter/pkg/analyzer"
	"log/slog"

	"go.uber.org/zap"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
	logger, _ := zap.NewProduction()
	slog.Info("Server started")
	slog.InfoContext(context.Background(), "Server start")
	logger.Info("server started!")
}
