package jobs

import "context"

type Job interface {
	// Run 返回错误则会重新调度
	Run(ctx context.Context) error
}
