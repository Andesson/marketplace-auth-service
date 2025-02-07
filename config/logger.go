package config

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	*log.Logger
	file *os.File
}

func NewLogger(logFilePath string) *Logger {
	// Criação do diretório de logs se não existir
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			log.Fatalf("❌ Falha ao criar diretório de logs: %v", err)
		}
	}

	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("❌ Falha ao abrir arquivo de log: %v", err)
	}

	return &Logger{
		Logger: log.New(logFile, "", log.LstdFlags),
		file:   logFile,
	}
}

func (l *Logger) flush() {
	if l.file != nil {
		if err := l.file.Sync(); err != nil {
			fmt.Println("Error flushing log file:", err)
		} else {
			fmt.Println("Log file flushed successfully.")
		}
	}
}

func (l *Logger) Debug(v ...interface{})   { l.Println(v...); l.flush() }
func (l *Logger) Info(v ...interface{})    { l.Println(v...); l.flush() }
func (l *Logger) Warning(v ...interface{}) { l.Println(v...); l.flush() }
func (l *Logger) Error(v ...interface{})   { l.Println(v...); l.flush() }

func (l *Logger) Debugf(format string, v ...interface{})   { l.Printf(format, v...); l.flush() }
func (l *Logger) Infof(format string, v ...interface{})    { l.Printf(format, v...); l.flush() }
func (l *Logger) Warningf(format string, v ...interface{}) { l.Printf(format, v...); l.flush() }
func (l *Logger) Errorf(format string, v ...interface{})   { l.Printf(format, v...); l.flush() }

func (l *Logger) Close() {
	if l.file != nil {
		if err := l.file.Close(); err != nil {
			fmt.Println("Error closing log file:", err)
		} else {
			fmt.Println("Log file closed successfully.")
		}
	}
}
