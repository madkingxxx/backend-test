package commands

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/madkingxxx/backend-test/internal/utils"
	"go.uber.org/zap"
)

type commandI interface {
	GetCronExpression() string
	GetName() string
	Run(ctx context.Context)
}

type Scheduler struct {
	scheduler *gocron.Scheduler
	commands  []commandI
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		scheduler: gocron.NewScheduler(time.UTC),
	}
}

func (s *Scheduler) Register(command commandI) {
	s.commands = append(s.commands, command)
}

func (s *Scheduler) warmUp(ctx context.Context) {
	for _, command := range s.commands {
		utils.Logger.Info(ctx, "warming up command:", zap.String("cron_name", command.GetName()))
		command.Run(ctx)
	}
}

func (s *Scheduler) Start(ctx context.Context) {
	utils.Logger.Info(ctx, "starting warm up...")
	s.warmUp(ctx)

	for _, command := range s.commands {
		if command.GetCronExpression() == "" {
			panic("cron expression is empty")
		}

		utils.Logger.Info(ctx, "scheduling command:", zap.String("cron_name", command.GetName()))
		_, err := s.scheduler.Cron(command.GetCronExpression()).Do(command.Run, ctx)
		if err != nil {
			utils.Logger.Error(ctx, "failed to schedule command:", zap.Error(err))
		}
	}

	utils.Logger.Info(ctx, "starting scheduler...")
	s.scheduler.StartAsync()
}
