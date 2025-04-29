package logger

import (
	"context"
	"log/slog"
	"os"
)

type MockLogger struct {
	logger *slog.Logger
}

func NewMockLogger() *MockLogger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	log := slog.New(handler)
	return &MockLogger{logger: log}
}

func (l *MockLogger) Debug(msg string, args ...any) {
	l.logger.DebugContext(context.Background(), msg, args...)
}

func (l *MockLogger) Info(msg string, args ...any) {
	l.logger.InfoContext(context.Background(), msg, args...)
}

func (l *MockLogger) Warn(msg string, args ...any) {
	l.logger.WarnContext(context.Background(), msg, args...)
}

func (l *MockLogger) Error(msg string, args ...any) {
	l.logger.ErrorContext(context.Background(), msg, args...)
}

func (l *MockLogger) With(args ...any) Logger {
	return &MockLogger{logger: l.logger.With(args...)}
}

func (l *MockLogger) WithContext(ctx context.Context) Logger {
	// A forma correta é criar um novo logger *sem* associar o contexto diretamente aqui.
	// A associação e extração de valores do contexto são feitas através do Handler's Attrer.
	// Retornamos uma nova instância do MockLogger com o logger existente.
	// A presença do contexto será tratada pelo Handler quando os logs forem formatados.
	return &MockLogger{logger: l.logger}
}

func (l *MockLogger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.logger.DebugContext(ctx, msg, args...)
}

func (l *MockLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

func (l *MockLogger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.logger.WarnContext(ctx, msg, args...)
}

func (l *MockLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
}
