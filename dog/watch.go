package dog

import (
	"fmt"
	"sccwatchdog/logger"
	"sccwatchdog/model"
	"sccwatchdog/utils"
	"strings"
	"time"
)

var log = logger.Logger()

func splitImageInfo(imageline string) []model.Image {
	images := strings.Split(imageline, " ")
	res := []model.Image{}
	var name, version string
	for _, image := range images {
		sImage := strings.Split(image, ":")
		i := sImage[0]
		si := strings.Split(i, "/")
		if len(si) > 1 {
			name = si[len(si)-1]
			version = sImage[1]
		} else {
			name = "NotGetName"
			version = "NotGetVersion"
		}
		res = append(res, model.Image{
			Name:    name,
			Version: version,
		})
	}
	return res
}

func changedInfo(deploy, beforeImage, lastImage, lastUpdateTime string) string {
	before := splitImageInfo(beforeImage)
	last := splitImageInfo(lastImage)
	changeCombine := ""
	for i, image := range before {
		changeCombine += fmt.Sprintf("%s:", image.Name)
		if image.Version != last[i].Version {
			changeCombine += fmt.Sprintf("%s ==> %s ", image.Version, last[i].Version)
		} else {
			changeCombine += fmt.Sprintf("%s ", image.Version)
		}
	}
	return fmt.Sprintf("%s has changed, %sat %v", deploy, changeCombine, lastUpdateTime)
}

func restartedInfo(deploy, image, lastUpdateTime, message string) string {
	before := splitImageInfo(image)
	restartCombine := ""
	for _, i := range before {
		restartCombine += fmt.Sprintf("%s:%s ", i.Name, i.Version)
	}
	return fmt.Sprintf("%s has restarted, %sat %v,%s", deploy, restartCombine, lastUpdateTime, message)
}

func watchDeployments(deploy *model.Deployment, s int, chi chan model.GetInfo) {
	for {
		time.Sleep(time.Duration(s) * time.Second)

		image, time, err := utils.GetDeploymentImageAndLastUpdateTime(deploy.Namespace, deploy.Name)

		if image != deploy.Image {
			chi <- model.GetInfo{
				Info:  changedInfo(deploy.Name, deploy.Image, image, time),
				Error: err,
			}
			deploy.Image = image
			deploy.LastUpdateTime = time
		} else if time != deploy.LastUpdateTime {
			message, _ := utils.GetdeploymentMessage(deploy.Namespace, deploy.Name)
			chi <- model.GetInfo{
				Info:  restartedInfo(deploy.Name, deploy.Image, time, message),
				Error: err,
			}
			deploy.LastUpdateTime = time
		}
		// log.Infof("%s : %s at %v", deploy.Name, image, time)
	}
}

func watchStart(deploys []model.Deployment, s int) {
	chi := make(chan model.GetInfo, len(deploys))
	for _, deploy := range deploys {
		go watchDeployments(&deploy, s, chi)
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

func getDeploymentInfoMap(deploys []model.Deployment) map[string]model.Deployment {
	deploysInfos := map[string]model.Deployment{}
	for _, deploy := range deploys {
		deploysInfos[deploy.Name] = deploy
	}
	return deploysInfos
}

func watchAllDeployments(deploysInfos map[string]model.Deployment, namespace string, s int, chi chan model.GetInfo) {

	for {
		time.Sleep(time.Duration(s) * time.Second)
		newDeploys, err := utils.GetAllDeploymentsInfos(namespace)
		for _, deploy := range newDeploys {
			if deploysInfos[deploy.Name].Name == "" {
				continue
			}
			if deploysInfos[deploy.Name].Image != deploy.Image {
				chi <- model.GetInfo{
					Info:  changedInfo(deploy.Name, deploysInfos[deploy.Name].Image, deploy.Image, deploy.LastUpdateTime),
					Error: err,
				}
				deploysInfos[deploy.Name] = deploy
			} else if deploysInfos[deploy.Name].LastUpdateTime != deploy.LastUpdateTime {
				message, _ := utils.GetdeploymentMessage(deploy.Namespace, deploy.Name)
				chi <- model.GetInfo{
					Info:  restartedInfo(deploy.Name, deploy.Image, deploy.LastUpdateTime, message),
					Error: err,
				}
				deploysInfos[deploy.Name] = deploy
			}
			// log.Infof("%s : %s at %v", deploy.Name, deploy.Image, deploy.LastUpdateTime)
		}
	}
}

func watchAllStart(deploys []model.Deployment, namespace, expect string, s int) {
	chi := make(chan model.GetInfo, len(deploys))
	deploysInfos := getDeploymentInfoMap(deploys)
	for _, eDeploy := range strings.Split(expect, " ") {
		deploysInfos[eDeploy] = model.Deployment{}
	}
	go watchAllDeployments(deploysInfos, namespace, s, chi)
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

func WatchDeplolyments(namespace, deployment, except string, s int) {
	realDeploys := []string{}
	if except != "" {
		for _, d := range strings.Split(deployment, " ") {
			if !strings.Contains(except, d) {
				realDeploys = append(realDeploys, d)
			}
		}
	} else {
		realDeploys = strings.Split(deployment, " ")
	}
	deploys, err := utils.GetDeploymentInfos(namespace, realDeploys)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	watchStart(deploys, s)
}

func WatchAllDeplolyments(namespace, except string, s int) {
	deploys, err := utils.GetAllDeploymentsInfos(namespace)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	watchAllStart(deploys, namespace, except, s)
}
