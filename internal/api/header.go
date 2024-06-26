package api

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type HeadersConfig struct {
	IdempotencyKey   string `mapstructure:"idempotency_key" yaml:"idempotency_key" default:"Idempotency-Key"`
	IdempotencyScope string `mapstructure:"idempotency_scope" yaml:"idempotency_scope" default:"Idempotency-Scope"`
}

func SupportedHeaders(cfg HeadersConfig) map[string]bool {
	return map[string]bool{
		cfg.IdempotencyKey:   true,
		cfg.IdempotencyScope: true,
	}
}

func GetHeaderString(ctx context.Context, headerKey string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	ikeys := md.Get(headerKey)
	if len(ikeys) < 1 {
		return ""
	}
	return ikeys[0]
}
