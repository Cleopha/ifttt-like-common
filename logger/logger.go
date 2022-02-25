package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type Mode uint8

const (
	Dev  Mode = iota
	Prod Mode = iota
)

func InitLogger(folder, logfile string, mode Mode) error {
	// Create the logs' directory if it does not exist.
	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory's path: %w", err)
	}

	if _, err = os.Stat(path + "/" + folder); os.IsNotExist(err) {
		if err = os.Mkdir(folder, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create log's directory: %w", err)
		}
	}

	// Creates the log file if it does not exist.
	file, err := os.OpenFile(path+"/"+folder+"/"+logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to create / open log file: %w", err)
	}

	var cfg zap.Config
	var encoderCfg zapcore.EncoderConfig

	switch mode {
	case Dev:
		cfg = zap.NewDevelopmentConfig()
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	case Prod:
		cfg = zap.NewProductionConfig()
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	// Creates the production logger and updates the output targets.
	cfg.OutputPaths = append(cfg.OutputPaths, file.Name())
	cfg.ErrorOutputPaths = append(cfg.ErrorOutputPaths, file.Name())

	// Updates the encoder configuration
	encoderCfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(time.RFC3339)
	}

	cfg.EncoderConfig = encoderCfg

	// Builds the logger based on the given configuration.
	production, err := cfg.Build()
	if err != nil {
		return fmt.Errorf("failed to build zap's production logger: %w", err)
	}

	zap.ReplaceGlobals(production)

	return nil
}
