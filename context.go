package context

import (
	"context"
	"time"
)

// ctx, cancel, err := context.WithDeadline(tkinfo.WaitingSeatTimeoutDuration)
func WithDeadline(dstring string) (context.Context, context.CancelFunc, error) {
	var duration time.Duration
	duration, err := time.ParseDuration(dstring)
	if err != nil {
		return nil, nil, err
	}
	d := time.Now().Add(duration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	return ctx, cancel, nil
}
