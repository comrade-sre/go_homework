package log

import "go.uber.org/zap"

func NewLogger(LOGPATH string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		LOGPATH,
	}
	return cfg.Build()
}
