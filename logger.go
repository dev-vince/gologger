package gologger

import (
	"fmt"
	"github.com/fatih/color"
)

type Level int

const (
	DebugLevel Level = 1
	InfoLevel Level = 2
	WarnLevel Level = 3
	ErrorLevel Level = 4
	SuccessLevel Level = 5
)

var useColor = true
var format = "%s: %s"
var loggerLevel = DebugLevel

func UseColor(color bool) {
	useColor = color
}

func (level Level) GetText() string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case SuccessLevel:
		return "SUCCESS"
	}
	return ""
}

func (level Level) GetColor() *color.Color {
	switch level {
	case DebugLevel:
		return color.New(color.FgWhite)
	case InfoLevel:
		return color.New(color.FgBlue)
	case WarnLevel:
		return color.New(color.FgYellow)
	case ErrorLevel:
		return color.New(color.FgRed)
	case SuccessLevel:
		return color.New(color.FgGreen)
	}
	return color.New(color.FgWhite)
}

func (level Level) Log(message string) {
	if level < loggerLevel {
		return
	}

	if useColor {
		fmt.Printf(format, level.GetColor().Sprint(level.GetText()), message)
	} else {
		fmt.Printf(format, level.GetText(), message)
	}
}

func Debug(message string) {
	DebugLevel.Log(message)
}

func Info(message string) {
	InfoLevel.Log(message)
}

func Warn(message string) {
	WarnLevel.Log(message)
}

func Error(message string) {
	ErrorLevel.Log(message)
}

func Success(message string) {
	SuccessLevel.Log(message)
}