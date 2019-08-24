package runtime

import "context"

func Context(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return ctx
}
