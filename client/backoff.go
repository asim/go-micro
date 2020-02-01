package client

import (
	"context"
	"github.com/micro/go-micro/v2/util/backoff"
	"time"
)

type BackoffFunc func(ctx context.Context, req Request, attempts int) (time.Duration, error)

// exponential backoff is a function x^e multiplied by a factor of 0.1 second.
func exponentialBackoff(ctx context.Context, req Request, attempts int) (time.Duration, error) {
	return backoff.Do(attempts), nil
}
