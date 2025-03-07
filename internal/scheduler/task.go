package scheduler

import (
	"go_server/pkg/util/cronscheduler"
	"go_server/pkg/util/log"
)

var scheduler *cronscheduler.Scheduler

func init() {
	scheduler = cronscheduler.NewScheduler()

}

// StartScheduler 启动定时任务调度器
func StartScheduler() {
	// UpdateMachineMapTask()
	scheduler.Start()
	log.Info("定时任务启动成功")
}

// func UpdateMachineMapTask() {
// 	services.UpdateMachineMap()
// 	scheduler.AddJob("*/10 * * * * *", func() {
// 		services.UpdateMachineMap()
// 	})
// }
