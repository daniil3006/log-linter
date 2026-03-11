package a

import (
	"fmt"
	"log/slog"
)

var log slog.Logger

func testLowercaseSlogSuccess() {
	log.Info("start server")
	slog.Info(fmt.Sprintf("fail start service on port %d", 8080))
	slog.Error("failed to connect to database")
}

func testLowercaseSlogFail() {
	log.Info("Starting server")     // want "lowercase"
	slog.Error("Failed to connect") // want "lowercase"
}

func testLatinSlogFail() {
	log.Info("cервер запустился")                                   // want "english"
	slog.Error("сервер" + "error")                                  // want "english"
	slog.Error(fmt.Sprintf("ошибка подключения на порту %d", 8080)) // want "english"
}

func testSpecialCharsSlogFail() {
	log.Info("start service!")                      // want "special characters"
	slog.Error("connection failed ☹️")              // want "special characters"
	slog.Warn("something" + "went" + "wrong" + "!") // want "special characters"
}

func testSensitiveDataSlogSuccess() {
	log.Info("token is valid")
	slog.Debug("api request completed")
}

func testSensitiveDataSlogFail() {
	token := "abcde"
	log.Info("token" + token)                     // want "sensitive"
	slog.Debug(fmt.Sprintf("api_key %s", "abcd")) // want "sensitive"
}

func testMultipleViolationsSlog() {
	log.Info("Success connect!") // want "lowercase" "special characters"
	slog.Error("Ошибка!")        // want "english" "special characters"
}
