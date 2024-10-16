package chans

import "context"

func Send[T any](ctx context.Context, dest chan<- T, value T) {
	select {
	case <-ctx.Done():
	case dest <- value:
	}
	return
}
