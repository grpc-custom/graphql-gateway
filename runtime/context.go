package runtime

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type authKey struct{}

const authorization = "authorization"

func Context(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	token := GetAuthToken(ctx)
	return metadata.AppendToOutgoingContext(ctx, authorization, token)
}

func SetAuthToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, authKey{}, token)
}

func GetAuthToken(ctx context.Context) string {
	value := ctx.Value(authKey{})
	token, ok := value.(string)
	if !ok {
		return ""
	}
	return token
}
