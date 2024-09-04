package server

import (
	"github.com/robfig/cron"
)

func (apiServer *apiServer) RunCronJob() {
	cron := cron.New()

	// Create cronjob
	//workers := utils.AppConfig.Workers
	//for _, worker := range workers {
	//	if worker.Enable {
	//		cronName := strings.TrimSpace(worker.Name)
	//		cronFunc := apiServer.mappingWorkers()[worker.Function]
	//		cronTime := strings.TrimSpace(worker.TimeRun)
	//		cron.AddFunc(cronTime, cronFunc)
	//		utils.ShowInfoLogs(fmt.Sprintf("Init worker [%s], time [%s] successful", cronName, cronTime))
	//	}
	//}

	cron.Start()
}

func (apiServer *apiServer) mappingWorkers() map[string]func() {
	mapHandler := make(map[string]func())

	return mapHandler
}
