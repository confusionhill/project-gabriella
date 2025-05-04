package utils

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2/widget"
)

var (
	LoggerInstance *Logger
)

type Logger struct {
	logs *widget.Label
}

func CreateLogger(logs *widget.Label) {
	LoggerInstance = &Logger{
		logs: logs,
	}
}

func (l *Logger) Log(message string) {
	l.logs.SetText(fmt.Sprintf("%s\n%s", l.logs.Text, message))
}

func (l *Logger) Debug(message string) {
	l.logs.SetText(fmt.Sprintf("[DEBUG] %s\n%s", l.logs.Text, message))
}

func (l *Logger) Error(message string) {
	l.logs.SetText(fmt.Sprintf("[ERROR] %s\n%s", l.logs.Text, message))
}

func (l *Logger) Info(message string) {
	l.logs.SetText(fmt.Sprintf("[INFO] %s\n%s", l.logs.Text, message))
}

func (l *Logger) Warn(message string) {
	l.logs.SetText(fmt.Sprintf("[WARN] %s\n%s", l.logs.Text, message))
}

func (l *Logger) Fatal(message string) {
	l.logs.SetText(fmt.Sprintf("[FATAL] %s\n%s", l.logs.Text, message))
	os.Exit(1)
}
