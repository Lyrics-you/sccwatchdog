package dog

import (
	"fmt"
	"sccwatchdog/utils"
	"strings"
)

func ShowDeployments(namespace, deployment string) {
	deploys := strings.Split(deployment, " ")
	deploy, err := utils.GetDeploymentInfos(namespace, deploys)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	for _, d := range deploy {
		fmt.Printf("[%s/%s: (%s) @ %s]\n", d.Namespace, d.Name, d.Image, d.LastUpdateTime)
	}
}

func ShowAllDeployments(namespace string) {
	deploy, err := utils.GetAllDeploymentsInfos(namespace)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	for _, d := range deploy {
		fmt.Printf("[%s/%s: (%s) @ %s]\n", d.Namespace, d.Name, d.Image, d.LastUpdateTime)
	}
}
