package cronservice

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CronServiceImpl struct {
	sugar         *zap.SugaredLogger
	CronScheduler *gocron.Scheduler
}

func NewCronJobService(sugar *zap.SugaredLogger) CronService {
	s := gocron.NewScheduler(time.UTC)
	s.TagsUnique()
	defer s.StartAsync()

	return &CronServiceImpl{
		sugar:         sugar,
		CronScheduler: s,
	}
}

func (c *CronServiceImpl) CreateRepetitionSecondJob(ctx context.Context, period int, task interface{}) {
	tag := "RepetitionTask-" + uuid.New().String()
	_, err := c.CronScheduler.
		Every(period).
		Second().
		Tag(tag).
		Do(task, ctx)

	if err != nil {
		c.sugar.Errorf("error invoke repetition job:%v", err.Error())
	}
}
