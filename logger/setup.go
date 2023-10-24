package logger

import (
	"fmt"

	"github.com/afshintalebi/go-gin-starter/config"
	"github.com/afshintalebi/go-gin-starter/utils"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
 * TODO: find a way to capture all messages that comes from Error & DPanic
 * https://stackoverflow.com/questions/70489167/uber-zap-logger-how-to-prepend-every-log-entry-with-a-string
 */

// must put `defer logger.Sync()` statement in the main function
var logger *zap.Logger

func getDevLogger() *zap.Logger {
	loggerRef := zap.NewDevelopmentEncoderConfig()
	loggerRef.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(loggerRef),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	))

	return logger
}

func getProdLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return logger

}

func init() {
	IsProductionMode := config.IsProductionMode()

	if IsProductionMode {
		logger = getProdLogger()
	} else {
		logger = getDevLogger()
	}

	logMessage := fmt.Sprintf("Logger has been loaded in %s mode", utils.GetEnvironmentText())
	fmt.Println(logMessage)
}
