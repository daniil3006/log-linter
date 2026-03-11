package a

import (
	"fmt"

	"go.uber.org/zap"
)

var logger = &zap.Logger{}

func testLowercaseZapSuccess() {
	logger.Info("start server")
	logger.Info(fmt.Sprintf("fail start service on port %d", 8080))
	logger.Error("failed to connect to database")
}

func testLowercaseZapFail() {
	logger.Info("Starting server")    // want "lowercase"
	logger.Error("FAILED TO CONNECT") // want "lowercase"
}

func testLatinZapFail() {
	logger.Info("cервер запустился")                                  // want "english"
	logger.Error("сервер" + "error")                                  // want "english"
	logger.Error(fmt.Sprintf("ошибка подключения на порту %d", 8080)) // want "english"
}

func testSpecialCharsZapFail() {
	logger.Info("start service!")                     // want "special characters"
	logger.Error("connection failed ☹️")              // want "special characters"
	logger.Warn("something" + "went" + "wrong" + "!") // want "special characters"
}

func testSensitiveDataZapSuccess() {
	logger.Info("token is valid")
	logger.Debug("api request completed")
}

func testSensitiveDataZapFail() {
	token := "abcde"

	logger.Info("token" + token)                    // want "sensitive"
	logger.Debug(fmt.Sprintf("api_key %s", "abcd")) // want "sensitive"
}

func testMultipleViolationsZap() {
	logger.Info("Success connect!") // want "lowercase" "special characters"
	logger.Error("Ошибка!")         // want "english" "special characters"
}
