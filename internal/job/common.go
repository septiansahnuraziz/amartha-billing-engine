package job

import (
	"amartha-billing-engine/config"
	"amartha-billing-engine/internal/entity"
	"amartha-billing-engine/internal/job/dto"
	"amartha-billing-engine/utils"
	"context"
	"time"

	"github.com/hibiken/asynq"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var mux = asynq.NewServeMux()

type (
	cronSpecAndRetention struct {
		CronSpec  string
		Retention time.Duration
		Timeout   time.Duration
	}

	bulkIndexFn func(ctx context.Context, IDs []uint) error
)

const (
	cronEvery1Hour     = "0 * * * *"
	cronEvery0030      = "30 00 * * *"
	cronEvery1Minute   = "* * * * *"
	cronEvery15Minutes = "*/15 * * * *"
	cronEvery5Minutes  = "*/5 * * * *"
)

func getParsedData[T any](ctx *context.Context, task *asynq.Task) (item T, err error) {
	var emptyValue T

	payload := dto.PayloadQueue{}
	if err := utils.JSONUnmarshal(task.Payload(), &payload); err != nil {
		return emptyValue, err
	}

	// set traceID
	*ctx = utils.SetTraceID(*ctx, payload.TraceID)

	return utils.MapToStruct[T](payload.Data)
}

func getCrontab(taskName entity.Task) string {
	return periodicJobsCronSpec[taskName].CronSpec
}

// AsynqTaskTracerMiddleware tracer for asynq task, place this middleware on mux
func taskTracerMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		tracer := otel.GetTracerProvider().Tracer(utils.MyCaller(1))

		ctx, span := tracer.Start(ctx, utils.WriteStringTemplate("[%s] WORKER task:%s", config.EnvironmentMode(), t.Type()))
		span.SetAttributes(attribute.String("task.type", t.Type()))
		span.SetAttributes(attribute.String("task.payload", string(t.Payload())))
		defer span.End()

		return h.ProcessTask(ctx, t)
	})
}
