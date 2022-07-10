package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"sccwatchdog/logger"
	"sccwatchdog/model"
	"strings"
	"sync"
)

var (
	log                = logger.Logger()
	ErrNamespaceEmpty  = errors.New("namespace is empty")
	ErrDeploymentEmpty = errors.New("deployment is empty")
)

func KubeCommand(exCommand []string, arg ...string) (string, error) {
	cmd := exec.Command("kubectl", exCommand...)
	// fmt.Println(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		// log.Errorf("%s, %s", err, strings.TrimSpace(string(out)))
		return "", errors.New(fmt.Sprintf("%s, %s", err, strings.TrimSpace(string(out))))
	}
	if strings.HasPrefix(string(out), "Error") {
		return "", errors.New(string(out))
	}
	return strings.TrimSpace(string(out)), nil
}

func GetAlldeploymentName(namespace string) ([]string, error) {
	if namespace == "" {
		return []string{}, ErrNamespaceEmpty
	}
	line := fmt.Sprintf(`get deployment -n %s -o jsonpath="%s"`, namespace, ItemsDeploymentName)
	cmdsplit := strings.Split(line, " ")
	out, err := KubeCommand(cmdsplit)
	if err != nil {
		return []string{}, err
	}
	names := strings.Split(out, " ")
	for id := range names {
		names[id] = ExtractString(names[id])
	}
	if len(names) == 1 && names[0] == "" {
		line := fmt.Sprintf(`get deployment -n %s`, namespace)
		out, _ := KubeCommand(strings.Split(line, " "))
		return []string{}, fmt.Errorf("%s", out)
	}
	return names, nil
}

func getDeploymentInfoByJsonPath(namespace, deployment, jsonpath string, chi chan model.GetInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	if deployment == "" {
		chi <- model.GetInfo{
			Info:  "",
			Error: ErrDeploymentEmpty,
		}
		return
	}
	line := fmt.Sprintf(`get deployment %s -n %s -o jsonpath="%s"`, deployment, namespace, jsonpath)
	cmdsplit := strings.Split(line, " ")
	out, err := KubeCommand(cmdsplit)
	if err != nil {
		chi <- model.GetInfo{
			Info:  "",
			Error: err,
		}
		return
	}
	info := ExtractString(out)
	chi <- model.GetInfo{
		Info:  info,
		Error: nil,
	}
	return
}

func GetdeploymentImage(namespace, deployment string) (string, error) {
	chi := make(chan model.GetInfo, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go getDeploymentInfoByJsonPath(namespace, deployment, DeploymentImage, chi, &wg)
	wg.Wait()
	info := <-chi
	if info.Error != nil {
		return "", info.Error
	}
	return info.Info, nil
}

func GetDeploymentContainers(namespace, deployment string) ([]string, error) {
	chi := make(chan model.GetInfo, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go getDeploymentInfoByJsonPath(namespace, deployment, DeploymentContainerName, chi, &wg)
	wg.Wait()
	info := <-chi
	if info.Error != nil {
		return []string{}, info.Error
	}
	containers := strings.Split(info.Info, " ")
	return containers, nil
}

func GetdeploymentLastUpdateTime(namespace, deployment string) (string, error) {
	chi := make(chan model.GetInfo, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go getDeploymentInfoByJsonPath(namespace, deployment, DeploymentLastUpdateTime, chi, &wg)
	wg.Wait()
	info := <-chi
	if info.Error != nil {
		return "", info.Error
	}
	time, err := TransTime(info.Info)
	if err != nil {
		return "", err
	}
	return time, nil
}

func GetDeploymentImageAndLastUpdateTime(namespace, deployment string) (string, string, error) {
	chi := make(chan model.GetInfo, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go getDeploymentInfoByJsonPath(namespace, deployment, DeploymentImageAndLastUpdateTime, chi, &wg)
	wg.Wait()
	info := <-chi
	if info.Error != nil {
		return "", "", info.Error
	}
	infosplit := strings.Split(info.Info, ",")
	image := infosplit[0]
	time, err := TransTime(infosplit[1])
	if err != nil {
		return "", "", err
	}
	return image, time, nil
}

func GetdeploymentMessage(namespace, deployment string) (string, error) {
	chi := make(chan model.GetInfo, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go getDeploymentInfoByJsonPath(namespace, deployment, DeploymentMessage, chi, &wg)
	wg.Wait()
	info := <-chi
	if info.Error != nil {
		return "", info.Error
	}
	return info.Info, nil
}

func GetDeploymentAllInfos(namespace string, chi chan model.GetInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	line := fmt.Sprintf(`get deployment -n %s -o jsonpath="%s"`, namespace, ItemsDeploymentAllInfos)
	cmdsplit := strings.Split(line, " ")
	out, err := KubeCommand(cmdsplit)
	if err != nil {
		chi <- model.GetInfo{
			Info:  "[]",
			Error: err,
		}
		return
	}
	imageAndtime := ExtractString(out)
	if imageAndtime == "" {
		chi <- model.GetInfo{
			Info:  "[]",
			Error: errors.New(fmt.Sprintf("No resources found in %s namespace", namespace)),
		}
		return
	}
	chi <- model.GetInfo{
		Info:  imageAndtime,
		Error: err,
	}
}

func GetAllDeploymentsInfos(namespace string) ([]model.Deployment, error) {
	if namespace == "" {
		namespace = "default"
	}
	chi := make(chan model.GetInfo, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go GetDeploymentAllInfos(namespace, chi, &wg)
	wg.Wait()
	info := <-chi
	if info.Error != nil {
		return []model.Deployment{}, info.Error
	}
	subInfo := info.Info[1 : len(info.Info)-1]
	infoSplit := strings.Split(subInfo, "][")

	deployments := []model.Deployment{}
	for _, inf := range infoSplit {
		si := strings.Split(inf, ",")
		time, _ := TransTime(si[2])
		deployments = append(deployments, model.Deployment{
			Name:           si[0],
			Namespace:      namespace,
			Image:          si[1],
			LastUpdateTime: time,
		})
	}
	return deployments, nil
}

func GetDeploymentInfos(namespace string, deploys []string) ([]model.Deployment, error) {
	if namespace == "" {
		namespace = "default"
	}
	deployments := []model.Deployment{}

	for _, deploy := range deploys {
		d := model.Deployment{
			Name:      deploy,
			Namespace: namespace,
		}
		image, time, _ := GetDeploymentImageAndLastUpdateTime(namespace, deploy)
		d.Image = image
		d.LastUpdateTime = time
		deployments = append(deployments, d)
	}
	return deployments, nil
}
