package cronservice

import "context"

type CronService interface {
	CreateRepetitionSecondJob(ctx context.Context, period int, task interface{})
}
