package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

type LogDetails struct {
	Operation string
	Function  string
	Message   string
}

func InitLog() {
	// for logging filename and line number
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func GetLogDetails(operation string, function string, message string) LogDetails {
	return LogDetails{
		Operation: operation,
		Function:  function,
		Message:   message,
	}
}

func LogInfo(ctx context.Context, details LogDetails) {
	logrus.WithFields(logrus.Fields{"operation": details.Operation, "function": details.Function, "txnId": ctx.Value("txnId")}).Info(details.Message)
}

func LogError(ctx context.Context, details LogDetails) {
	logrus.WithFields(logrus.Fields{"operation": details.Operation, "function": details.Function, "txnId": ctx.Value("txnId")}).Error(details.Message)
}

func LogPanic(ctx context.Context, details LogDetails) {
	logrus.WithFields(logrus.Fields{"operation": details.Operation, "function": details.Function, "txnId": ctx.Value("txnId")}).Panic(details.Message)
}

func LogFatalAndStop(ctx context.Context, details LogDetails) {
	logrus.WithFields(logrus.Fields{"operation": details.Operation, "function": details.Function, "txnId": ctx.Value("txnId")}).Fatal(details.Message)
}
