package policies

import "context"

type Policy interface {
	Authorize(ctx context.Context) error
}
