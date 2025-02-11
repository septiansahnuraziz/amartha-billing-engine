package job

import (
	"amartha-billing-engine/config"
	"amartha-billing-engine/internal/entity"
	"amartha-billing-engine/utils"
	"errors"
	"time"

	"github.com/hibiken/asynq"
	log "github.com/sirupsen/logrus"
)

type TaskProcessor struct {
	server      *asynq.Server    // for processing task
	scheduler   *asynq.Scheduler // for scheduling
	taskHandler *TaskHandler
	namespace   string
}

// ConstructProcessor use in console job
func ConstructProcessor(redisOpt asynq.RedisConnOpt, timeLocation *time.Location, namespace string, taskHandler *TaskHandler, taskQueue entity.TaskQueue) *TaskProcessor {
	logger := log.WithField("job-namespace", namespace)

	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Concurrency: config.WorkerConcurrency(),
			Queues: map[string]int{
				namespace: 10,
			},
			Logger: logger,
			//ErrorHandler: asynq.ErrorHandlerFunc(processorErrorHandler(taskQueue)),
		},
	)

	scheduler := asynq.NewScheduler(redisOpt, &asynq.SchedulerOpts{
		Logger:   logger,
		Location: timeLocation,
		PreEnqueueFunc: func(t *asynq.Task, opts []asynq.Option) {
			taskName := t.Type()
			crontab := getCrontab(entity.Task(taskName))
			opts = append(opts, asynq.TaskID(utils.WriteStringTemplate("%s_%s", taskName, utils.GetCronNextAt(crontab))))
			newTask := asynq.NewTask(taskName, nil, opts...)
			*t = *newTask
		},
		PostEnqueueFunc: func(info *asynq.TaskInfo, err error) {
			if err == nil {
				return
			}

			if errors.Is(err, asynq.ErrDuplicateTask) || errors.Is(err, asynq.ErrTaskIDConflict) {
				log.Info(err)
				return
			}

			log.Error(err)
		},
	})

	return &TaskProcessor{
		server:      server,
		taskHandler: taskHandler,
		scheduler:   scheduler,
		namespace:   namespace,
	}
}

// Run job
func (t *TaskProcessor) Run() {
	t.registerTasks()
	t.registerCronTask()

	go func() {
		if err := t.scheduler.Run(); err != nil {
			log.Fatalf("failed to run scheduler : %v", err)
		}
	}()

	if err := t.server.Start(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

// Stop to stop job client
func (t *TaskProcessor) Stop() {
	if t.server != nil {
		t.server.Stop()
		t.server.Shutdown()
	}
}

func (t *TaskProcessor) registerTasks() {
	mux.Use(taskTracerMiddleware)

	//mux.HandleFunc(TaskSendTelegramNotification.String(), t.taskHandler.sendTelegramOnError)
}

func (t *TaskProcessor) registerCronTask() {
	for taskName, cronSpecAndRetention := range periodicJobsCronSpec {
		if _, err := t.scheduler.Register(
			getCrontab(taskName),
			asynq.NewTask(taskName.String(), nil),
			asynq.Queue(t.namespace),
			asynq.Retention(cronSpecAndRetention.Retention),
			asynq.MaxRetry(config.WorkerRetryAttempts()),
			asynq.Timeout(cronSpecAndRetention.Timeout),
		); err != nil {
			log.Fatal(err)
		}
	}
}

// This variable is used to define the specifications of the registered cron
var periodicJobsCronSpec = map[entity.Task]cronSpecAndRetention{}
