package utils

import "context"

type contextKey string

const TraceIDKey contextKey = "trace_id"

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

func GetTraceID(ctx context.Context) string {
	val, ok := ctx.Value(TraceIDKey).(string)
	if !ok {
		return ""
	}
	return val
}