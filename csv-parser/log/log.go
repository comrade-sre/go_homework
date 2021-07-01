package log

import 	"go.uber.org/zap"
//const (
//    InfoLevel = zapcore.InfoLevel
//    ErrorLevel = zapcore.ErrorLevel
//    PanicLevel = zapcore.PanicLevel
//)
func NewLogger(LOGPATH string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		LOGPATH,
	}
	return cfg.Build()
}
