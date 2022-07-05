package dog

import (
	"fmt"
	"sccwatchdog/logger"
	"sccwatchdog/model"
	"sccwatchdog/utils"
	"time"
)

var log = logger.Logger()

func WatchDeployments(deploy model.Deployment, s int, chi chan model.GetInfo) {
	for {
		time.Sleep(time.Duration(s) * time.Second)

		image, time, err := utils.GetDeploymentImageAndLastUpdateTime(deploy.Namespace, deploy.Name)

		if image != deploy.Image {
			chi <- model.GetInfo{
				Info:  fmt.Sprintf("%s has changed,%s==>%s at %v", deploy.Name, deploy.Image, image, time),
				Error: err,
			}
			deploy.Image = image
			deploy.LastUpdateTime = time
		} else if time != deploy.LastUpdateTime {
			message, _ := utils.GetDepolymentMessage(deploy.Namespace, deploy.Name)
			chi <- model.GetInfo{
				Info:  fmt.Sprintf("%s has restarted,%s at %v,%s", deploy.Name, deploy.Image, time, message),
				Error: err,
			}
			deploy.LastUpdateTime = time
		}
		// log.Infof("%s : %s at %v", deploy.Name, image, time)
	}
}

func WatchStart(deploys []model.Deployment, s int) {
	chi := make(chan model.GetInfo, len(deploys))
	for _, deploy := range deploys {
		go WatchDeployments(deploy, s, chi)
	}
	for {
		select {
		case info := <-chi:
			if info.Error != nil {
				log.Warnf("%v", info.Error)
			} else {
				log.Info(info.Info)
			}
		}
	}
}

func WatchAllDeployments(deploys []model.Deployment, namespace string, s int, chi chan model.GetInfo) {
	deploysInfos := map[string]model.Deployment{}
	for _, deploy := range deploys {
		deploysInfos[deploy.Name] = deploy
	}
	for {
		time.Sleep(time.Duration(s) * time.Second)
		newDeploys, err := utils.GetAllDeploymentsInfos(namespace)
		for _, deploy := range newDeploys {
			if deploysInfos[deploy.Name].Image != deploy.Image {
				chi <- model.GetInfo{
					Info:  fmt.Sprintf("%s has changed,%s==>%s at %v", deploy.Name, deploysInfos[deploy.Name].Image, deploy.Image, deploy.LastUpdateTime),
					Error: err,
				}
				deploysInfos[deploy.Name] = deploy
			} else if deploysInfos[deploy.Name].LastUpdateTime != deploy.LastUpdateTime {
				message, _ := utils.GetDepolymentMessage(deploy.Namespace, deploy.Name)
				chi <- model.GetInfo{
					Info:  fmt.Sprintf("%s has restarted at %v,%s", deploy.Name, deploy.LastUpdateTime, message),
					Error: err,
				}
				deploysInfos[deploy.Name] = deploy
			}
			// log.Infof("%s : %s at %v", deploy.Name, deploy.Image, deploy.LastUpdateTime)
		}
	}
}

func WatchAllStart(deploys []model.Deployment, namespace string, s int) {
	chi := make(chan model.GetInfo, len(deploys))
	go WatchAllDeployments(deploys, namespace, s, chi)
	for {
		select {
		case info := <-chi:
			if info.Error != nil {
				log.Warnf("%v", info.Error)
			} else {
				log.Info(info.Info)
			}
		}
	}
}
