package logger

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mdobak/go-xerrors"
)

type stackFrame struct {
	Func   string `json:"func"`
	Source string `json:"source"`
	Line   int    `json:"line"`
}
type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger() *SlogLogger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		//ReplaceAttr: replaceAttr,
		//AddSource:   true,
	})
	log := slog.New(handler)
	return &SlogLogger{logger: log}
}

func (l *SlogLogger) Debug(msg string, args ...any) {
	l.logger.DebugContext(context.Background(), msg, args...)
}

func (l *SlogLogger) Info(msg string, args ...any) {
	l.logger.InfoContext(context.Background(), msg, args...)
}

func (l *SlogLogger) Warn(msg string, args ...any) {
	l.logger.WarnContext(context.Background(), msg, args...)
}

func (l *SlogLogger) Error(msg string, args ...any) {
	l.logger.ErrorContext(context.Background(), msg, args...)
}

func (l *SlogLogger) With(args ...any) Logger {
	return &SlogLogger{logger: l.logger.With(args...)}
}

func (l *SlogLogger) WithContext(ctx context.Context) Logger {
	// A forma correta é criar um novo logger *sem* associar o contexto diretamente aqui.
	// A associação e extração de valores do contexto são feitas através do Handler's Attrer.
	// Retornamos uma nova instância do SlogLogger com o logger existente.
	// A presença do contexto será tratada pelo Handler quando os logs forem formatados.
	return &SlogLogger{logger: l.logger}
}

func (l *SlogLogger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.logger.DebugContext(ctx, msg, args...)
}

func (l *SlogLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

func (l *SlogLogger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.logger.WarnContext(ctx, msg, args...)
}

func (l *SlogLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
}

func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	/* switch a.Value.Kind() {
	case slog.KindAny:
		switch v := a.Value.Any().(type) {
		case error:
			a.Value = fmtErr(v)
		}
	}
	return a */
	if a.Key == slog.SourceKey {
		source, ok := a.Value.Any().(slog.Source)
		if ok {
			// Se o nível do log for ERROR, manter apenas o nome do arquivo
			if slog.LevelError.String() == slog.LevelError.String() { // Sempre verdadeiro para simplificar o exemplo
				parts := strings.Split(source.File, "/")
				a.Value = slog.StringValue(parts[len(parts)-1])
			} else {
				// Para outros níveis, manter o formato padrão (arquivo:linha)
				a.Value = slog.StringValue(source.File + ":" + strconv.Itoa(source.Line))
			}
		}
	}
	return a
}

// marshalStack extracts stack frames from the error
func marshalStack(err error) []stackFrame {
	trace := xerrors.StackTrace(err)
	if len(trace) == 0 {
		return nil
	}
	frames := trace.Frames()
	s := make([]stackFrame, len(frames))
	for i, v := range frames {
		f := stackFrame{
			Source: filepath.Join(
				filepath.Base(filepath.Dir(v.File)),
				filepath.Base(v.File),
			),
			Func: filepath.Base(v.Function),
			Line: v.Line,
		}
		s[i] = f
	}
	return s
}

// fmtErr returns a slog.Value with keys `msg` and `trace`. If the error
// does not implement interface { StackTrace() errors.StackTrace }, the `trace`
// key is omitted.
func fmtErr(err error) slog.Value {
	var groupValues []slog.Attr
	groupValues = append(groupValues, slog.String("msg", err.Error()))
	frames := marshalStack(err)
	if frames != nil {
		groupValues = append(groupValues,
			slog.Any("trace", frames),
		)
	}
	return slog.GroupValue(groupValues...)
}
