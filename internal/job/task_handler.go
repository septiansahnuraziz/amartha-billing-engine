package job

import (
	"amartha-billing-engine/internal/entity"
	"time"

	"gorm.io/gorm"
)

type TaskHandler struct {
	jktTimeLoc *time.Location
	taskQueue  entity.TaskQueue
	db         *gorm.DB
}

func ConstructTaskHandler(
	jktTimeLoc *time.Location,
	db *gorm.DB,
	taskQueue entity.TaskQueue,
) *TaskHandler {
	return &TaskHandler{
		jktTimeLoc: jktTimeLoc,
		db:         db,
		taskQueue:  taskQueue,
	}
}

//func (th *TaskHandler) handleUpdateSalesAgentESCriteriaByChannelCashierMappingID(c context.Context, task *asynq.Task) error {
//	ctx, span := otel.StartWorker(c)
//	defer span.End()
//	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
//		"ctx":     utils.DumpIncomingContext(ctx),
//		"payload": utils.ByteToString(task.Payload()),
//	})
//
//	criteria, err := getParsedData[entity.SalesAgentUpdateCriteriaESEnqueuePayload](&ctx, task)
//	if err != nil {
//		logger.Error(err)
//		return err
//	}
//	return th.salesAgentService.UpdateSalesAgentCriteriaByChannelCashierMappingIdES(ctx, utils.ExpectedNumber[uint64](criteria.ChannelCashierMappingId), criteria.ToSalesAgentUpdateCriteria())
//}
