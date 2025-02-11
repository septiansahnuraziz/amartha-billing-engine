package job

import (
	"amartha-billing-engine/config"
	"amartha-billing-engine/internal/entity"
	"amartha-billing-engine/internal/job/dto"
	"amartha-billing-engine/utils"
	"context"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

type (
	taskQueue struct {
		client    *asynq.Client // for enqueueing jobs
		namespace string
	}
)

// ConstructQueue use for enqueueing jobs only
func ConstructQueue(redisOpt asynq.RedisConnOpt, namespace string) entity.TaskQueue {
	client := asynq.NewClient(redisOpt)

	return &taskQueue{
		client:    client,
		namespace: namespace,
	}
}

func (queue *taskQueue) Enqueue(ctx context.Context, taskName entity.Task, data any) error {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"ctx":  utils.DumpIncomingContext(ctx),
		"data": utils.Dump(data),
	})

	payload := dto.PayloadQueue{
		Data:    data,
		TraceID: utils.GetTraceID(ctx),
	}

	marshalled, err := utils.JSONMarshal(payload)
	if err != nil {
		logger.Error(err)
		return err
	}

	task := asynq.NewTask(taskName.String(), marshalled)
	if _, err := queue.client.Enqueue(task, queue.generalOpts()...); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (queue *taskQueue) Stop() {
	if queue.client != nil {
		_ = queue.client.Close()
	}
}

func (queue *taskQueue) generalOpts() []asynq.Option {
	return []asynq.Option{
		asynq.Queue(queue.namespace),
		asynq.Retention(config.WorkerTaskRetention()),
		asynq.MaxRetry(config.WorkerRetryAttempts()),
	}
}
