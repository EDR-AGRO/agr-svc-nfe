package logger

// zap logger
import (
	"br/com/agr/nfe/infrastructure/apm"
	"context"
	"fmt"

	"go.uber.org/zap"

	"go.elastic.co/apm/module/apmzap/v2"
)

var logger = zap.NewExample(zap.WrapCore((&apmzap.Core{}).WrapCore))

func Info(ctx context.Context, msg string) {
	traceContextFields := apmzap.TraceContext(ctx)
	logger.With(traceContextFields...).Info(msg)
}

func Infof(ctx context.Context, msg string, args ...interface{}) {
	traceContextFields := apmzap.TraceContext(ctx)
	msg = fmt.Sprintf(msg, args...)
	logger.With(traceContextFields...).Info(msg)
}

func Error(ctx context.Context, msg string) {
	traceContextFields := apmzap.TraceContext(ctx)
	logger.With(traceContextFields...).Error(msg)
}

func Errorf(ctx context.Context, msg string, args ...interface{}) {
	traceContextFields := apmzap.TraceContext(ctx)
	msg = fmt.Sprintf(msg, args...)
	logger.With(traceContextFields...).Error(msg)
}

func Debug(ctx context.Context, msg string) {
	traceContextFields := apmzap.TraceContext(ctx)
	logger.With(traceContextFields...).Debug(msg)
}

func Debugf(ctx context.Context, msg string, args ...interface{}) {
	traceContextFields := apmzap.TraceContext(ctx)
	msg = fmt.Sprintf(msg, args...)
	logger.With(traceContextFields...).Debug(msg)
}

func Warn(ctx context.Context, msg string) {
	traceContextFields := apmzap.TraceContext(ctx)
	logger.With(traceContextFields...).Warn(msg)
}

func Warnf(ctx context.Context, msg string, args ...interface{}) {
	traceContextFields := apmzap.TraceContext(ctx)
	msg = fmt.Sprintf(msg, args...)
	logger.With(traceContextFields...).Warn(msg)
}

func Fatal(ctx context.Context, msg string) {
	traceContextFields := apmzap.TraceContext(ctx)
	logger.With(traceContextFields...).Fatal(msg)
}

func Fatalf(ctx context.Context, msg string, args ...interface{}) {
	traceContextFields := apmzap.TraceContext(ctx)
	msg = fmt.Sprintf(msg, args...)
	logger.With(traceContextFields...).Fatal(msg)
}

func Panic(ctx context.Context, msg string) {
	defer apm.Flush()
	traceContextFields := apmzap.TraceContext(ctx)
	logger.With(traceContextFields...).Panic(msg)
}

func Panicf(ctx context.Context, msg string, args ...interface{}) {
	defer apm.Flush()
	traceContextFields := apmzap.TraceContext(ctx)
	msg = fmt.Sprintf(msg, args...)
	logger.With(traceContextFields...).Panic(msg)
}
